//Import the functions you need from the SDKs you need
/*import { initializeApp } from "firebase/app";
import { getStorage } from "firebase/storage";

// Your web app's Firebase configuration
const firebaseConfig = {
  apiKey: "AIzaSyC08wZzoDB-TojuqkdjM4S3h-TtGNLf7gI",
  authDomain: "prac-ec.firebaseapp.com",
  projectId: "prac-ec",
  storageBucket: "prac-ec.appspot.com",
  messagingSenderId: "216459314207",
  appId: "1:216459314207:web:55ad381c9131426971012a"
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);
// Initialize Cloud Storage and get a reference to the service

const storage = getStorage(app);

function CreateItem(){
      
    var files = document.querySelector("#photo_url").files;
    console.log(files);
    
    var image = files[0];
    console.log(image);
    
    var ref = storage.ref().child(image.name);
    console.log(ref);

    ref.put(image).then(function(snapshot) {
      alert('アップロードしました');
      console.log(snapshot);
    });
  }*/

/*const OnFileUploadToFirebase = (e) => {
    const file = e.target.files[0];
    const storageRef = ref(storage, "image/" + file.name);
    //uploadBytes(storageRef, file).then((snapshot) => {
    //    console.log('Uploaded a blob or file!');
    //  });
      const uploadImage = uploadBytesResumable(storageRef, file);
        uploadImage.on("state_changed",
         (snapshot) => {
            setLoading(true);
        },
        (error) => {
            console.log(error);
        },
        () => {
            setLoading(false);
            setUploaded(true);
         }
        );
};*/
