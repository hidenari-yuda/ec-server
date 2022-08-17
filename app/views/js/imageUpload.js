function OutputImage(target){
    var reader = new FileReader();
    
    reader.onload = function () {
      $samplePhotoURL = $("#samplePhotoURL");
      
      img = new Image();
      img.src = this.result;
      img.onload = function() {
        $samplePhotoURL.css("width", 500);
        $samplePhotoURL.css("height", 400);
      }
      
      $samplePhotoURL.css("background", "url(" + this.result + ") center center / contain no-repeat");
    }
    
    reader.readAsDataURL(target.files[0]);
  }