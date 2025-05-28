import { startClock } from "./src/js/clock.js";
import { environmentHandler } from "./src/js/environment.js";

import { connectGoserverJS } from "./src/js/rocketdatamain.js";
import { weatherApiDataHandler } from "./src/js/weather.js";

async function initializeComponents() {
  await loadComponent("clock", "./src/components/clock.html");
  await loadComponent("timer", "./src/components/timer.html");
  await loadComponent("rocketData", "./src/components/rocketData.html");
  await loadComponent("environment", "./src/components/environment.html");
  await loadComponent("weather", "./src/components/weather.html");
}

async function apiLoader() {
  window.api
    .fetchWeatherData()
    .then((data) => {
      console.log(data);
      localStorage.setItem("weatherApiObject", JSON.stringify(data));
    })
    .catch((error) => {
      console.error("Error fetching weather data:", error);
    });
}

async function loadComponent(elementId, filePath) {
  console.log(`Loading component: ${elementId} from ${filePath}`);
  var weatherApiResult;
  fetch(filePath)
    .then((response) => response.text())
    .then((data) => {
      const container = document.getElementById(elementId);
      if (!container) {
        console.error(`container with ${elementId} not found`);
        return;
      }
      // container.innerHTML = data;
      while(container.firstChild) 
        container.removeChild(container.firstChild)
      container.appendChild(document.createTextNode(data));
      
      if (elementId === "clock") {
        startClock(container);
      }
      if (elementId === "weather") {
        var storedData = localStorage.getItem("weatherApiObject");
        var parsedData = JSON.parse(storedData);
        const weatherElement = document.getElementById("weather");
        weatherApiDataHandler(weatherElement, parsedData);
      }
      if (elementId === "environment") {
        var storedData = localStorage.getItem("weatherApiObject");
        var parsedData = JSON.parse(storedData);
        const environmentElement = document.getElementById("environment");
        environmentHandler(environmentElement, parsedData);

      }

      if(elementId === "rocketData") {
        connectGoserverJS();
      }
    })
    .catch((error) => console.error("Error loading component:", error));
}

document.addEventListener("DOMContentLoaded", () => {
  apiLoader();
  initializeComponents();
});
