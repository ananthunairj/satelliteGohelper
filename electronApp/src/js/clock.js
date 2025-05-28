export function startClock(container) {
  function updateClock() {
    var clockelement = new HTMLElement();
    clockelement = container;

    const today = new Date();

    const hours = today.getHours();
    const minutes = today.getMinutes();
    const seconds = today.getSeconds();

    const hour = hours < 10 ? "0" + hours : hours;
    const minute = minutes < 10 ? "0" + minutes : minutes;
    const second = seconds < 10 ? "0" + seconds : seconds;

    const hourTime = hour > 12 ? hour - 12 : hour;

    const ampm = hour < 12 ? "AM" : "PM";

    const time = hourTime + ":" + minute + ":" + second + " " + ampm;
    const dateTimeElement = clockelement.querySelector("#date-time");
    
    if (dateTimeElement) {
      // dateTimeElement.innerHTML = time;
      while(dateTimeElement.firstChild) 
        dateTimeElement.removeChild(dateTimeElement.firstChild);
      dateTimeElement.appendChild(document.createTextNode(time));
    } else {
      console.error("time element not found.");
    }
    setTimeout(updateClock, 1000);
  }
  updateClock();
}