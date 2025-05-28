import { LinkedListVariables } from "./constgroups.js";
import { DoublyLinkedListJS } from "./linkedfordoublemain.js";



export async function connectGoserverJS() {
  let rawEndpoints = window.env.API_ENDPOINTS;
  try {
    if (!rawEndpoints) return console.error("Endpoint in env corrupted..");
    let apiEndpoints = JSON.parse(rawEndpoints);
    const endpoint = apiEndpoints.goservice;
    const dataInlinklist = new DoublyLinkedListJS();
    dataInlinklist = await getGODataJs(endpoint);
    // const string45 = JSON.stringify(dataInlinklist)
    // console.log(string45)

  } catch (e) {
    console.error("Failed to parse API_ENDPOINTS:", e);
  }
}

async function getGODataJs(args) {
  try {
    const endpoint = `http://localhost:8080/${args}`;
    const eventSource = new EventSource(endpoint);
    const resultcontainer = new DoublyLinkedListJS();
    eventSource.addEventListener("message", function (event) {
      const data = JSON.parse(event.data);
      if (data.Flag === true) {
        console.log(data);
        resultcontainer.insertAtEnd(data.Data);

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

async function arrayConvertforPlot(linkedlistobj) {
  try {
    let timeArray = await linkedListToArray(linkedlistobj,LinkedListVariables.TIME);
    let velocityArray = await linkedListToArray(linkedlistobj,LinkedListVariables.VELOCITY);
    let angleArray = await linkedListToArray(linkedlistobj, LinkedListVariables.ANGLE);
   

  } catch (e){
        console.error(`Error occured ${e}`);
  }
}

async function linkedListToArray(linkedlistobj, key) {
  const convertinglinklist = new DoublyLinkedListJS();
  convertinglinklist = linkedlistobj;
  let size = convertinglinklist.sizereturner();
  const result = [size]
  for (let node = convertinglinklist.head; node !== null; node = node.next) {
    if (node.data && key in node.data) {
      result.push(node.data[key])
    } else {
      result.push(null);
    }
  }
  return result;
}



