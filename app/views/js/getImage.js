/*import { getStorage, ref, uploadBytes } from "firebase/storage";


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
    // Get a reference to the storage service, which is used to create references in your storage bucket
    //const storage = getStorage();
    
    // Create a storage reference from our storage service
    //const storageRef = ref(storage);
    
    // Create a child reference
    //const imagesRef = ref(storage, 'images');
    // imagesRef now points to 'images'
    
    // Child references can also take paths delimited by '/'
    //const spaceRef = ref(storage, 'images/space.jpg');
    // spaceRef now points to "images/space.jpg"
    // imagesRef still points to "images"
}*/