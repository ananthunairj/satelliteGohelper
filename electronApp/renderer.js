import { startClock } from "./src/js/clock.js";
import { environmentHandler } from "./src/js/environment.js";

import { connectGoserverJS, loadPlotlyIfNeeded } from "./src/js/rocketdatamain.js";
import { weatherApiDataHandler } from "./src/js/weather.js";


document.addEventListener("DOMContentLoaded",async () => {
  await apiLoader();
  await initializeComponents();
});

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

async function initializeComponents() {
  await loadComponent("clock", "./src/components/clock.html");
  await loadComponent("environment", "./src/components/environment.html");
  await loadComponent("weather", "./src/components/weather.html");
   await loadComponent("rocketData", "./src/components/rocketData.html");
}



async function loadComponent(elementId, filePath) {
  console.log(`Loading component: ${elementId} from ${filePath}`);

  try {
    const response = await fetch(filePath);
    const data = await response.text();

    const container = document.getElementById(elementId);
    if (!container) {
      console.error(`Container with ID "${elementId}" not found`);
      return;
    }

    container.innerHTML = data;

    if (elementId === "clock") {
      startClock(container);
    }

    if (elementId === "weather") {
      const storedData = localStorage.getItem("weatherApiObject");
      const parsedData = JSON.parse(storedData);
      const weatherElement = document.getElementById("weather");
      weatherApiDataHandler(weatherElement, parsedData);
    }

    if (elementId === "environment") {
      const storedData = localStorage.getItem("weatherApiObject");
      const parsedData = JSON.parse(storedData);
      const environmentElement = document.getElementById("environment");
      environmentHandler(environmentElement, parsedData);
    }

    if (elementId === "rocketData") {
      await loadPlotlyIfNeeded();
      await connectGoserverJS(container); 
    }
  } catch (error) {
    console.error("Error loading component:", error);
  }
}

