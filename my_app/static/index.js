let wrapper = document.getElementById("wrapper_data");
let cartWrapper = document.getElementById("cart_wrapper");
let $ = document.querySelector.bind(document)
let $$ = document.querySelectorAll.bind(document)
let productsData = []
let productsElements = []
let existIncart = []
// Make a Fetch API request to get product data

async function fetchData(){
    await fetch("/products")
    .then(response => response.json())
    .then(data => {
      // Loop through each product in the data
      console.log(data)
      productsData = data
      existIncart = Array(data.length).fill(false)
      console.log(existIncart)
      data.forEach(product => {
        // Create a new product element
        let productElement = document.createElement("div");
        productElement.className = "body_product";
  
        // Populate the product element with data
        productElement.innerHTML = `
          <div class="body_product_img" style="background-color: rgb(225, 231, 237);">
            <img src="${product.image}">
          </div>
          <div class="body_product_name">${product.name}</div>
          <div class="body_product_description">${product.description}</div>
          <div class="body_product_bottom">
            <div class="body_product_bottom_price">${product.price}</div>
            <div class="item_addcart_btn">
              <p>ADD TO CART</p>
              <img width="20" height="20" src="https://img.icons8.com/ios-filled/20/checkmark--v1.png" alt="checkmark--v1"/>
            </div>
           
          </div>
        `;
        // Append the product element to the wrapper
        wrapper.appendChild(productElement);
        productsElements.push(productElement)
      });
    })
    .catch(error => {
      console.error("Error fetching product data:", error);
    });
}



  function addToCardHandler(){
    var products = $$(".body_product")
    products.forEach((element,index) => {
        element.querySelector('.item_addcart_btn img').hidden = true
        element.querySelector('.item_addcart_btn').onclick = () =>{
          

            addToCart(index)
        }
    });
  }
  function addToCart(position){

    if(existIncart[position]){
        return
    }
    existIncart[position]++ 

    let productElement = document.createElement("div");
    productElement.className = "cart_item";
    let product = productsData[position]
    // Populate the product element with data
    productElement.innerHTML = `
                <div class="cart_item_left cartItemLeft">
                <div class="cart_item_img cartItemImage"
                    style="background-color: rgb(34, 175, 220);">
                    <div class="cart_item_img_item"><img
                            src="${product.image}">
                    </div>
                </div>
            </div>
            <div class="App_cartItemRight_2LNcC cartItemRight">
                <div class="cart_item_name cartItemName">${product.name}</div>
                <div class="cart_item_price cartItemPrice">$${product.price}</div>
                <div class="cart_item_actions cartItemActions">
                    <div class="cart_item_count cartItemCount">
                        <div class="cart_item_count_btn cart_item_count_subtract">-</div>
                        <div class="cart_item_count_amount">1</div>
                        <div class="cart_item_count_btn cart_item_count_add">+</div>
                    </div>
                    <div class="cart_item_remove cartItemRemove"><img
                            src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAEAAAABACAYAAACqaXHeAAAABHNCSVQICAgIfAhkiAAAAAlwSFlzAAAB2AAAAdgB+lymcgAAABl0RVh0U29mdHdhcmUAd3d3Lmlua3NjYXBlLm9yZ5vuPBoAAALISURBVHic7Zs9bxNBEIYfgyUKAhhQUhDRICEBCh0fgkhBNIT8gPwZ6Gig5y8QCUhH5AqE3EZJgQRKEDSpKEAQkTMdcijGRvi8Z+/e3eze4X2kKe40t/Pu+LRfN4bIdNNQbLsJ3ATOFWznC7AJ/C6syCMngC3gsCTb7LdZGx5SXucH9kBD6BGNRoGrNWlTLQEa7R5VaFMtAbXBZwLWkVnHxtZ9iZr6N6Bp6TcHXAOOW/qfz7i36un5X8A28NXSfywrQJfypzVtS4D7ZSRgpwKdyWsfJnXOZincxf7VrxoJcHKcg80g2ClFShg6ZTQyD2xQr3GgC7yi+EYs8t+TZ329gKwJfiLzbRU4Cywh/fmuGegpw/PssmYwS5aAfURTD3ikFegKo4PNe61gDrxjWFMPuGj7sMte4JLh3mWH57VYSF03cDg7cEmAabxQ2aM7UkjX1O8GfSRgHmgjM8YO4wfOFWC379umYguZVcyrrkm0U/4JMGvwm2N0tblh0b5Jk+222csbcCd1PYOsI9KYzhvuqij6Bx8JMO0kZyz91HehcRAMLSA0MQGhBYQmJiC0gNDEBIQWEJqYgNACQhMTEFpAaGICQgsITUxAaAGhiQnwEMP0+axr6af+6c1HAjqp6wQpo02zxWhi3moIykveU+FBfUGCfEq7N8Z3GSlrSbD/vl/oVNiFvAnQpvLH4pUmJsDBN2tEDlnHn1UBZppljLgkYC/j/i2HNspmMeP+nkawY8ABowPOa41gFjSQaTKt5wDRqsKaIeAh8Bjd/x+laQBPMrQ80wy8iJSgmAK/QWpzW4rxW8gndNMvPyiPua0YH4DnGcGrYGuK/f7LGeBjgM5Nsl3gtGK/h7gAfFbukIt96mvySgt4WVB4UesBL4BTyn0dy42+iEGxog/bR8ai60XFlzl1NZFiyllknNDgB/ANKbaq1V9pI1XlD82w8ru3YIVHAAAAAElFTkSuQmCC">
                    </div>
                </div>
            </div>
    `;
    // Append the product element to the wrapper
    cartWrapper.appendChild(productElement);
    productElement.querySelector(".cart_item_count_add").onclick = () =>{
       let num = parseInt(productElement.querySelector(".cart_item_count_amount").innerText)
        num++
        existIncart[position] ++
        productElement.querySelector(".cart_item_count_amount").innerText = num
        updateTotalAmount()

    }
    productElement.querySelector(".cart_item_count_subtract").onclick = () =>{
        
        let num = parseInt(productElement.querySelector(".cart_item_count_amount").innerText)
         num--
         existIncart[position] --
         productElement.querySelector(".cart_item_count_amount").innerText = num
         updateTotalAmount()
         if(existIncart[position] == 0){
            productElement.remove()
         }
     }
     productElement.querySelector(".cartItemRemove").onclick = () =>{
        
        existIncart[position] = 0;
        productElement.remove()

         updateTotalAmount()
     }
    updateTotalAmount()
  }

  async function runApp(){


    await fetchData()
    addToCardHandler()
  }
  function updateTotalAmount(){
    let sum = 0;
    for(let i =0;i<productsData.length;i++){
        sum+= productsData[i].price*existIncart[i]
        let element =  productsElements[i]
        if(existIncart[i] > 0){
            element.querySelector('.item_addcart_btn img').hidden = false
            element.querySelector('.item_addcart_btn p').hidden = true
        }
        else {
            element.querySelector('.item_addcart_btn img').hidden = true
            element.querySelector('.item_addcart_btn p').hidden = false
        }

    }
    $(".cart_cont_title_sum").innerText = (Math.round(sum * 100) / 100).toFixed(2);


    if(sum == 0)
        $(".card_empty_txt").hidden = false
    else  $(".card_empty_txt").hidden = true
    
  }
  runApp()
  