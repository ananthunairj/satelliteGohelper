import windThreshold from "../utils/thresholds.js";


export async function environmentHandler(environmentElement, data) {
  windSpeedHandler(environmentElement, data);
  windDirectionFinder(environmentElement, data);
}

async function windSpeedHandler(environmentElement, data) {
  var environelement = new HTMLElement();
  environelement = environmentElement;
  var windData = data.hourly;
  let windSpeedKmh =
    windData.wind_speed_10m[windData.wind_speed_10m.length - 1];

  var windSpeedms = Number(((windSpeedKmh * 5) / 18).toFixed(1));
  const environmentDisplay =
    environelement.querySelector("#environmentdata");
  if (environmentDisplay) {
    while(environmentDisplay.firstChild)
      environmentDisplay.removeChild(environmentDisplay.firstChild);
    environmentDisplay.appendChild(document.createTextNode("🍃 " + windSpeedms + " m/s"))
    // environmentDisplay.innerHTML = "🍃 " + windSpeedms + " m/s";
  } else {
    console.log("id weather missing");
  }
  if (windSpeedms >= windThreshold) {
    environmentDisplay.insertAdjacentHTML("beforeend", "❌");
  } else {
    environmentDisplay.insertAdjacentHTML("beforeend", "  ✅");
  }
}

async function windDirectionFinder(environmentElement, data) {
    var windelement = new HTMLElement();
  windelement = environmentElement;
  
  var windData = data.hourly;
  var directionWind;
  let windDirection =
    windData.wind_direction_10m[windData.wind_direction_10m.length - 1];

  switch (true) {
    case windDirection <= 180:
      if (windDirection > 90) {
        if (windDirection <= 134) {
          directionWind = "South East";
        } else {
          directionWind = "South";
        }
      } else {
        if (windDirection == 0) {
          directionWind = "North";
        } else if (windDirection <= 44) {
          directionWind = "North East";
        } else {
          directionWind = "East";
        }
      }
      break;

    case windDirection > 180:
      if (windDirection <= 224) {
        directionWind = "South West";
      } else if (windDirection <= 270) {
        directionWind = "West";
      } else if (windDirection <= 314) {
        directionWind = "North West";
      } else {
        directionWind = "North";
      }
      break;

    default:
      console.log("Error occured in calculation");
  }
  let allowedDirections = ["South East", "North East", "East"];
  let indicatorElement;
  let directionIndicator = windDirectionSafetyChecker(
    allowedDirections,
    directionWind
  );
  if (directionIndicator) indicatorElement = "✅";
  else indicatorElement = "⚠️";

  const windDisplay = windelement.querySelector("#windDirection");
  if (windDisplay) {
    // windDisplay.innerHTML = "🧭 " + directionWind + " " + indicatorElement;
    while(windDisplay.firstChild) 
      windDisplay.removeChild(windDisplay.firstChild)
    windDisplay.appendChild(document.createTextNode("🧭 " + directionWind + " " + indicatorElement));
  } else {
    console.log("direction wind missing");
  }


}

function windDirectionSafetyChecker(allowedDirections, directionWind) {
  let indicator = false;
  for (let element of allowedDirections) {
    if (element == directionWind) {
      indicator = true;
      break;
    }
  }
  return indicator;
}


