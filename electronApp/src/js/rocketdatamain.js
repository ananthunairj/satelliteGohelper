import { DoublyLinkedListJS } from "./linkedfordoublemain.js";

async function getGODataJs(args) {
  try {
    const endpoint = `http://localhost:8080/${args}`;
    const eventSource = new EventSource(endpoint);
    const resultcontainer = new DoublyLinkedListJS();
    eventSource.addEventListener("message", function (event) {
      const data = JSON.parse(event.data);
      if (data.Flag === true) {
        console.log(data);
        resultcontainer.insertAtBeginning(data.Data);
        
      } else if (data.Flag === false) {
        console.log("closing connection");
        eventSource.close();
        return resultcontainer;
      }
    });
  } catch (error) {
    console.error(`Error: ${error}`);
  }
}

export async function connectGoserverJS() {
  let rawEndpoints = window.env.API_ENDPOINTS;
  try {
    if (!rawEndpoints) return console.error("Endpoint in env corrupted..");
    let apiEndpoints = JSON.parse(rawEndpoints);
    const endpoint = apiEndpoints.goservice;
    const dataInlinklist = new DoublyLinkedListJS();
    dataInlinklist = getGODataJs(endpoint);
    // const string45 = JSON.stringify(dataInlinklist)
    // console.log(string45)

  } catch (e) {
    console.error("Failed to parse API_ENDPOINTS:", e);
  }
}
