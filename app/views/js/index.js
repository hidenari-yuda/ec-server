import { uploadBytes } from "firebase/storage";

function body() {
    var dates = document.getElementsByName("created");

    var len = dates.length;

    for (var i = 0; i < len; i++){
      var datesText = dates[i].innerText
      subDates = datesText.substring(0, 11);

      dates[i].innerText = subDates
    }

    //var fileReader = new FileReader();

      var items_images = document.getElementsByClassName("items_images");
      console.log(items_images);
      
      var lens = items_images.length;
      console.log(lens);

      
      for (var i = 0; i < lens; i++){

      var src = items_images[i].getAttribute("src");
      console.log(src)

      // var srcURL = fileReader.readAsDataURL(src[i]);
       var srcURL = window.URL.createObjectURL(src[i]);

      console.log(srcURL)

    }

  }

  function OutputImage(target){
    var reader = new FileReader();

    reader.onload = function () {
        $samplePhotoURL = $("#samplePhotoURL");

        img = new Image();
        img.src = this.result;
        console.log(this.result);
        img.onload = function() {
            $samplePhotoURL.css("width", 500);
            $samplePhotoURL.css("height", 400);
        }

        $samplePhotoURL.css("background", "url(" + this.result + ") center center / contain no-repeat");
    }

    reader.readAsDataURL(target.files[0]);
}

function CreateItem(){

  var Inputimg = document.querySelector("#photo_url");
  var src = Inputimg.getAttribute("src");

  
  uploadBytes(src, file).then((snapshot) => {
      console.log('Uploaded a blob or file!');
  });

function tododelete(){ 
    const result = confirm('削除しますか');
    if(result){
      console.log('削除しました');
    }else{
      console.log('削除をとりやめました');
    }
}
}


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