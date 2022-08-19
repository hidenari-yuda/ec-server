document.addEventListener('DOMContentLoaded', function() {
    var dates = document.getElementsByName("created");

    var len = dates.length;

    for (var i = 0; i < len; i++){
      var datesText = dates[i].innerText
      subDates = datesText.substring(0, 11);

      dates[i].innerText = subDates
    }
    
  })

 function confirmDelete(){ 
    const result = confirm('削除しますか');

    if (result){
      alert('削除しました');
    } else{
      alert('削除をとりやめました');
    }
  }

  //var fileReader = new FileReader();

  /*var items_images = document.getElementsByClassName("items_images");
    console.log(items_images);
    
    var lens = items_images.length;
    console.log(lens);

    
    for (var i = 0; i < lens; i++){

    var src = items_images[i].getAttribute("src");
    console.log(src)

    // var srcURL = fileReader.readAsDataURL(src[i]);
     var srcURL = window.URL.createObjectURL(src[i]);

    console.log(srcURL)

  }*/
  
  
  /*var datesText = new Date(dates[i].innerText)
  console.log(datesText)
  const year = datesText.getFullYear()
  const month = datesText.getMonth() + 1
  const day = datesText.getDate()
  
  var subDates = `${year}/${month}/${day}`
  
  // createdAtを表示形式変更
  /*function dateFormat() {
    const elements = document.getElementsByClassName(`todo-created`);
    Array.prototype.filter.call(elements, function (el) {
      const resDate = new Date(el.innerText)
      const year = resDate.getFullYear()
      const month = resDate.getMonth() + 1
      const day = resDate.getDate()
      const dateString = `${year}/${month}/${day}`
      el.textContent = dateString
    })
  }*/