const { contextBridge, ipcRenderer } = require("electron");

contextBridge.exposeInMainWorld("electron", {
  ipcRenderer: ipcRenderer,
});

contextBridge.exposeInMainWorld("api", {
  fetchWeatherData: () => ipcRenderer.invoke("fetch-weather-data"),
});


contextBridge.exposeInMainWorld('electronAPI', {
   runGoScript: (args) => ipcRenderer.invoke('run-go-script', args),
  });
  