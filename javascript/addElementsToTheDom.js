class Food {
  constructor(root, data) {
    this.root = root;
    this.data = data;
  }

  renderList() {
    // event listener is added to the root b/c of event delegation
    // this.root.addEventListener(e => {
    //   const { target } = e;
    //   target.remove();
    // });
    const fragment = document.createDocumentFragment();
    this.data.forEach(i => {
      fragment.appendChild(this.createElem(i));
    });
    this.root.appendChild(fragment);
  }

  createElem(item) {
    const itemElem = document.createElement("li");
    itemElem.textContent = item;
    return itemElem;
  }
}
let rootElement = document.getElementById("ul");
let foodData = ["🍔", "🍣"];

const foods = new Food(rootElement, foodData);
foods.renderList();

// <html>
//   <body>
//     <ul id="ul"></ul>
//   </body>
// </html>

// Code should output a list of emojis.
