const { app, BrowserWindow, ipcMain } = require("electron");
const path = require("node:path");
const dotenv = require("dotenv");
const axios = require("axios");
const fetch = require('node-fetch');


dotenv.config();

const createWindow = () => {
  const win = new BrowserWindow({
    width: 800,
    height: 600,
    fullscreen: true,

    webPreferences: {
      preload: path.join(__dirname, "preload.js"),
      nodeIntegration: true,
      contextIsolation: true,
    },
  });

  win.loadFile("index.html");
};

app.whenReady().then(() => {
  createWindow();
});

ipcMain.handle("fetch-weather-data", async () => {
  const launchSpotData = JSON.parse(process.env.LAUNCH_SPOT_DATA);
  const latitudeData = launchSpotData.latitude;
  const longitudeData = launchSpotData.longitude;
  console.log(launchSpotData);

  const url = "https://api.open-meteo.com/v1/forecast";

  const params = {
    latitude: latitudeData,
    longitude: longitudeData,
    hourly: "temperature_2m,weather_code,wind_speed_10m,wind_direction_10m",
    hourly_units: "temperature_2m",
  };

  try {
    const response = await axios.get(url, { params });
    return response.data;
  } catch (error) {
    console.error(error);
    throw error;
  }
});

app.on("window-all-closed", () => {
  if (process.platform !== "darwin") app.quit();
});

