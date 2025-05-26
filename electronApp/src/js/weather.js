import weatherData from "../utils/weatherDictionary.js";

export function weatherApiDataHandler(weatherElement, data) {
  var utilWeather = data.hourly;
  let checkerWeather = utilWeather.weather_code[utilWeather.weather_code.length - 1];
  let temperature = utilWeather.temperature_2m[utilWeather.temperature_2m.length - 1];

  let weatherMapper = weatherData[checkerWeather];

  const weatherDisplay = weatherElement.querySelector("#weatherdata");
  if (weatherDisplay) {
    weatherDisplay.innerHTML = weatherMapper + " " + "🌡️ " + temperature +" " + data.hourly_units.temperature_2m;
  } else {
    console.log("id weather missing");
  }
}
