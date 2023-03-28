document.addEventListener("DOMContentLoaded", () => {
    const buttons = document.querySelectorAll(".moving-button");
  
    buttons.forEach((button) => {
      button.addEventListener("mouseover", moveButton);
    });
  
    function moveButton(event) {
      const screenWidth = window.innerWidth;
      const screenHeight = window.innerHeight;
      const buttonWidth = event.target.offsetWidth;
      const buttonHeight = event.target.offsetHeight;
      const x = Math.floor(Math.random() * (screenWidth - buttonWidth));
      const y = Math.floor(Math.random() * (screenHeight - buttonHeight));
  
      event.target.style.transition = "transform 1s ease-in-out";
      event.target.style.transform = `translate(${x}px, ${y}px)`;
  
      event.target.style.backgroundColor = randomColor();
      event.target.style.borderRadius = `${Math.floor(Math.random() * 50)}%`;
      
    }
  
    function randomColor() {
      const red = Math.floor(Math.random() * 256);
      const green = Math.floor(Math.random() * 256);
      const blue = Math.floor(Math.random() * 256);
      return `rgb(${red}, ${green}, ${blue})`;
    }
  
    setInterval(() => {
      buttons.forEach((button) => {
        button.style.transition = "none";
        button.style.transform = "none";
      });
    }, 5000);
  });
  