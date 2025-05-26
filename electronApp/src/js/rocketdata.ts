
import { RocketPositiondata, StimulationResult } from "./interfaces";
import { DoublyLinkedLIst } from "./linkedfordouble";

declare global {
  interface Window {
    env : {
      API_ENDPOINTS : string;
    }
  }
}

async function getGoData(args: string) {
  try {
    const endpoint: string = `http://localhost:8080/${args}`;
    const eventSource: EventSource = new EventSource(endpoint);
    const resultcontainer = DoublyLinkedLIst.createLinkedlist<RocketPositiondata>();
    eventSource.addEventListener ("message",function (event) {
      const data: StimulationResult = JSON.parse(event.data);
      if (data.Flag === true) {
         console.log(data);
        resultcontainer.insertAtBeginning(data.Data);
        console.log("Current data is " ,resultcontainer.traverseBackward());
      } else if (data.Flag === false)  {
        console.log("closing connection");
        eventSource.close();
       
      }
    });
    
  } catch (error) {
    console.error(`Error: ${error}`);
  }
}

export async function connectGoserver() {
  let rawEndpoints = window.env.API_ENDPOINTS;
  type ApiEndpoints = {
    goservice: string;
    [key: string]: string;
  };
  try {
    if (!rawEndpoints) return console.error("Endpoint in env corrupted..");
    let apiEndpoints: ApiEndpoints = JSON.parse(rawEndpoints);
    const endpoint: string = apiEndpoints.goservice;
    getGoData(endpoint);
  } catch (e) {
    console.error("Failed to parse API_ENDPOINTS:", e);
  }
}
