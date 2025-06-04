import { LinkedListVariables } from "./constgroups.js";
import { DoublyLinkedListJS } from "./linkedfordoublemain.js";

export async function connectGoserverJS(container) {
  try {
    let rawEndpoints = window.env.API_ENDPOINTS;
    if (!rawEndpoints) return console.error("Endpoint in env corrupted..");
    let apiEndpoints = JSON.parse(rawEndpoints);
    let endpoint = apiEndpoints.goservice;

    const dataInlinklist = await getGODataJs(endpoint);
    await plotMaker(dataInlinklist, container);
    dataInlinklist.clear();
    // const string45 = JSON.stringify(dataInlinklist)
    // console.log(string45)
  } catch (e) {
    console.error("Failed to parse API_ENDPOINTS:", e);
  }
}

// async function getGODataJs(args) {
//   try {
//     const endpoint = `http://localhost:8080/${args}`;
//     const eventSource = new EventSource(endpoint);
//     const resultcontainer = new DoublyLinkedListJS();
//     eventSource.addEventListener("message", function (event) {
//       const data = JSON.parse(event.data);
//       if (data.Flag === true) {
//         console.log(data)
//         resultcontainer.insertAtEnd(data.Data);

//       } else if (data.Flag === false) {
//         console.log("closing connection");
//         eventSource.close();
//         size = resultcontainer.sizereturner()
//         return resultcontainer, size;
//       }
//     });
//   } catch (error) {
//     console.error(`Error: ${error}`);
//   }
// }

// async function getGODataJs(args) {
//   return new Promise((resolve, reject) => {
//     try {
//       const endpoint = `http://localhost:8080/${args}`;
//       const socket = new WebSocket(endpoint, { withCredentials: true });
//       const resultcontainer = new DoublyLinkedListJS();

//       eventSource.addEventListener("message", (event) => {
//         const data = JSON.parse(event.data);
//         if (data.Flag === true) {
//           resultcontainer.insertAtEnd(data.Data);
//         } else if (data.Flag === false) {
//           eventSource.close();
//            console.log("Received data length:", resultcontainer.length);
//           resolve(resultcontainer);
//         }
//       });

//       eventSource.addEventListener("error", (err) => {
//         eventSource.close();
//         reject(err);
//       });
//     } catch (err) {
//       reject(err);
//     }
//   });
// }
async function getGODataJs(args) {
  return new Promise((resolve, reject) => {
    try {
      let receivedCounts = new Set();
      const endpoint = `ws://localhost:8080/${args}`;
      const socket = new WebSocket(endpoint);
      const resultcontainer = new DoublyLinkedListJS();
      let count = 0;
      socket.binaryType = "arraybuffer";

      socket.onopen = () => {
        console.log("WebSocket connected");
      };

      socket.onmessage = (event) => {
        const text = new TextDecoder().decode(event.data);
        const data = JSON.parse(text);
        console.log("Received:", data);
        count++;

        if (data.Flag === true) {
          resultcontainer.insertAtEnd(data.Data);
          receivedCounts.add(data.Count);
        } else if (data.Flag === false) {
          for (let i = 1; i <= 973; i++) {
            if (!receivedCounts.has(i)) {
              console.log("Missing packet:", i);
            }
          }
          setTimeout(() => {
            socket.send("ACK");
            socket.close();
            resolve(resultcontainer);
          }, 100);
        }
      };

      socket.onerror = (err) => {
        console.error("WebSocket error:", err);
        reject(err);
      };

      socket.onclose = () => {
        console.log("WebSocket closed");
      };
      socket.onerror = (err) => console.error("WebSocket error:", err);
    } catch (err) {
      reject(err);
    }
  });
}

async function plotMaker(linkedlistobj, container) {
  try {
    const Plotly = window.Plotly;
    const velotimeDiv = container.querySelector("#velotime");
    const angletimeDiv = container.querySelector("#angletime");
    if (!velotimeDiv) {
      console.error("Plot div with ID 'velotimeDiv' not found in the DOM.");
      return;
    }
    if (!angletimeDiv) {
      console.error("Plot div with ID 'angletimeDiv' not found in the DOM.");
      return;
    }

    if (typeof window.Plotly === "undefined") {
      console.error(
        "Plotly is not available. Make sure it's loaded or exposed via preload.js."
      );
      return;
    }

    let timeArray = await linkedListToArray(
      linkedlistobj,
      LinkedListVariables.TIME
    );
    let velocityArray = await linkedListToArray(
      linkedlistobj,
      LinkedListVariables.VELOCITY
    );
    console.log(timeArray);
    let angleArray = await linkedListToArray(
      linkedlistobj,
      LinkedListVariables.ANGLE
    );

    var tracetimevelo = {
      x: timeArray,
      y: velocityArray,
      mode: "lines+markers",
      type: "scatter",
      line: {
        color: "royalblue",
        width: 2,
      },
      marker: {
        size: 4,
        color: "darkblue",
      },
      name: "Velocity vs Time",
      hovertemplate: "Time: %{x}s<br>Velocity: %{y} km/s<extra></extra>",
    };

    var layouttimevelo = {
      title: {
        text: "Time vs Velocity",
        font: {
          size: 24,
          family: "Arial, sans-serif",
        },
      },
      xaxis: {
        title: "Time (s)",
        showgrid: true,
        gridcolor: "#eee",
        zeroline: false,
        tickfont: { size: 12 },
        titlefont: { size: 16 },
      },
      yaxis: {
        title: "Velocity (km/s)",
        showgrid: true,
        gridcolor: "#eee",
        zeroline: false,
        tickfont: { size: 12 },
        titlefont: { size: 16 },
        range: [0, 10],
      },
      margin: {
        l: 60,
        r: 30,
        t: 60,
        b: 50,
      },
      plot_bgcolor: "#f9f9f9",
      paper_bgcolor: "#fff",
      shapes: [
        {
          type: "line",
          x0: 300,
          x1: 300,
          y0: 0,
          y1: 10,
          line: {
            color: "red",
            width: 2,
            dash: "dot",
          },
        },
      ],
    };

    var tracetimeangle = {
      x: timeArray,
      y: angleArray,
      mode: "lines+markers",
      type: "scatter",
      line: {
        color: "royalblue",
        width: 2,
      },
      marker: {
        size: 4,
        color: "darkblue",
      },
      name: "Time vs Angle",
      hovertemplate: "Time: %{x}s<br>Angle: %{y}° <extra></extra>",
    };

    var layouttimeangle = {
      title: {
        text: "Time vs Angle",
        font: {
          size: 24,
          family: "Arial, sans-serif",
        },
      },
      xaxis: {
        title: "Time (s)",
        showgrid: true,
        gridcolor: "#eee",
        zeroline: false,
        tickfont: { size: 12 },
        titlefont: { size: 16 },
      },
      yaxis: {
        title: "Angle (°)",
        showgrid: true,
        gridcolor: "#eee",
        zeroline: false,
        tickfont: { size: 12 },
        titlefont: { size: 16 },
        range: [90, 0],
      },
      margin: {
        l: 60,
        r: 30,
        t: 60,
        b: 50,
      },
      plot_bgcolor: "#f9f9f9",
      paper_bgcolor: "#fff",
      shapes: [
        {
          type: "line",
          x0: 300,
          x1: 300,
          y0: 0,
          y1: 10,
          line: {
            color: "red",
            width: 2,
            dash: "dot",
          },
        },
      ],
    };

    var config = {
      responsive: true,
      displayModeBar: true,
      displaylogo: false,
      modeBarButtonsToRemove: [
        "select2d",
        "lasso2d",
        "autoScale2d",
        "resetScale2d",
        "hoverClosestCartesian",
        "hoverCompareCartesian",
      ],
    };

    await Plotly.newPlot(velotimeDiv, [tracetimevelo], layouttimevelo, config);
    await Plotly.newPlot(
      angletimeDiv,
      [tracetimeangle],
      layouttimeangle,
      config
    );
    // await clearArray(timeArray,velocityArray,angleArray);
  } catch (e) {
    console.error(`Error occured ${e}`);
  }
}

async function clearArray(arrayone, arraytwo, arraythree) {
  arrayone.length = 0;
  arraytwo.length = 0;
  arraythree.length = 0;
}

// async function linkedListToArray(linkedlistobj, key) {
//   let convertinglinklist = new DoublyLinkedListJS();
//   convertinglinklist = linkedlistobj;
//   let result = [size]
//   for (let node = convertinglinklist.head; node !== null; node = node.next) {
//     if (node.data && key in node.data) {
//       result.push(node.data[key])
//     } else {
//       result.push(null);
//     }
//   }
//   return result;
// }

async function linkedListToArray(linkedlistobj, key) {
  const size = linkedlistobj.getsize();
  const result = [size];
  for (let node = linkedlistobj.head; node; node = node.next) {
    result.push(node.data?.[key] ?? null);
  }
  return result;
}

export async function loadPlotlyIfNeeded() {
  if (typeof window.Plotly !== "undefined") return;

  return new Promise((resolve, reject) => {
    const script = document.createElement("script");
    script.src = "src/libs/plotly.min.js";
    script.onload = () => {
      console.log("Plotly loaded dynamically.");
      resolve();
    };
    script.onerror = () => {
      reject(new Error("Failed to load Plotly"));
    };
    document.head.appendChild(script);
  });
}
