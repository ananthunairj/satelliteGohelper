const { contextBridge, ipcRenderer } = require("electron");
// const Plotly = require('plotly.js-dist-min');

contextBridge.exposeInMainWorld("electron", {
  ipcRenderer: ipcRenderer,
});

contextBridge.exposeInMainWorld("api", {
  fetchWeatherData: () => ipcRenderer.invoke("fetch-weather-data"),
});


contextBridge.exposeInMainWorld('electronAPI', {
   runGoScript: (args) => ipcRenderer.invoke('run-go-script', args),
  });

 contextBridge.exposeInMainWorld('env', {
  API_ENDPOINTS : process.env.API_ENDPOINTS,
 }) 

//  contextBridge.exposeInMainWorld('Plotly', Plotly);
