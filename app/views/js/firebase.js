// Import the functions you need from the SDKs you need
import { initializeApp } from "firebase/app";
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

/*<script type="module">
  // Import the functions you need from the SDKs you need
  import { initializeApp } from "https://www.gstatic.com/firebasejs/9.9.2/firebase-app.js";
  // TODO: Add SDKs for Firebase products that you want to use
  // https://firebase.google.com/docs/web/setup#available-libraries

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
</script>*/