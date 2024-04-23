<script>
  import { goto } from "$app/navigation";
  import {
    GoogleAuthProvider,
    GithubAuthProvider,
    signInWithPopup,
    getAuth,
    signOut,
  } from "firebase/auth";
  import { initializeApp } from "firebase/app";
  import { currentUser } from "../stores.js";

  // firebase constants
  const firebaseConfig = {
    apiKey: "AIzaSyDAT9rLWYjZ84ebly-thXKf23uxGKFRKV8",
    authDomain: "itec2024revamp.firebaseapp.com",
    projectId: "itec2024revamp",
    storageBucket: "itec2024revamp.appspot.com",
    messagingSenderId: "620055731881",
    appId: "1:620055731881:web:c170cf6563be7c065153b8",
    measurementId: "G-GTQR42DLSZ",
  };
  const app = initializeApp(firebaseConfig);
  const googleProvider = new GoogleAuthProvider();
  const githubProvider = new GithubAuthProvider();
  const auth = getAuth();

  let currentUserUID = "";
  currentUser.subscribe((value) => {
    currentUserUID = value;
  });

  function loginRequest(email, username, uid) {
    const user = {
      email: email,
      username: username,
      uid: uid,
    };
    fetch("http://127.0.0.1:3000/addUser", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(user),
    })
      .then((response) => {
        console.log(response);
        response.json();
      })
      .catch((error) => {
        console.error("There was a problem with the fetch operation:", error);
      });
  }

  // authentication functions
  function loginWithGoogle() {
    signInWithPopup(auth, googleProvider)
      .then((result) => {
        // This gives you a Google Access Token. You can use it to access the Google API.
        const credential = GoogleAuthProvider.credentialFromResult(result);
        const token = credential.accessToken;

        const user = result.user;
        loginRequest(user.email, user.displayName, user.uid);
        currentUser.set(user.uid);
      })
      .catch((error) => {
        const errorCode = error.code;
        const errorMessage = error.message;
        const email = error.customData.email;
        const credential = GoogleAuthProvider.credentialFromError(error);
        console.log("error logging in with google: ");
        console.log(errorMessage);
        // FIX: currently alerts regardless of the error
        alert(
          "An account with this email already exists. Try logging in with Github."
        );
      });
  }

  function loginWithGitHub() {
    signInWithPopup(auth, githubProvider)
      .then((result) => {
        // This gives you a Google Access Token. You can use it to access the Google API.
        const credential = GithubAuthProvider.credentialFromResult(result);
        const token = credential.accessToken;

        const user = result.user;
        currentUser.set(user.uid);
        console.log(user + "from github");
      })
      .catch((error) => {
        const errorCode = error.code;
        const errorMessage = error.message;
        const email = error.customData.email;
        const credential = GithubAuthProvider.credentialFromError(error);
        // FIX: currently alerts regardless of the error
        alert(
          "An account with this email already exists. Try logging in with Google."
        );
        console.log("error logging in with github: ");
        console.log(errorMessage);
      });
  }

  function logout() {
    signOut(auth)
      .then(() => {
        currentUser.set("");
        console.log("signed out");
      })
      .catch((error) => {
        console.log(error);
        console.log("error signing out");
      });
  }

  let selectedPublic = false;
  function goToPublicApps() {
    selectedDev = false;
    selectedPublic = !selectedPublic;
    goto("/publicApps");
  }

  let selectedDev = false;
  function goToDevDashboard() {
    selectedPublic = false;
    selectedDev = !selectedDev;
    goto("../developerDashboard");
  }
</script>

<nav class="navbar navbar-expand-lg bg-primary" style="padding: 18px 32px">
  <div class="container-fluid">
    <h1 class="navbar-child" style="margin-bottom: 0">ItecMonitor</h1>
    {#if currentUserUID != ""}
      <button type="button" class="btn btn-primary login-bttn" on:click={logout}
        >Logout</button
      >
    {:else}
      <div class="dropdown">
        <button
          type="button"
          class="btn dropdown-toggle login-bttn"
          data-bs-toggle="dropdown"
          data-bs-display="static"
          aria-expanded="false"
          style="color: white"
        >
          Login
        </button>
        <ul class="dropdown-menu dropdown-menu-lg-end">
          <li>
            <button class="login-options" on:click={loginWithGoogle}
              >Login with Google</button
            >
          </li>
          <li>
            <button class="login-options" on:click={loginWithGitHub}
              >Login with Github</button
            >
          </li>
        </ul>
      </div>
    {/if}
  </div>
</nav>
<div class="content-container">
  <div class="nav-buttons-container">
    <button
      class="{selectedPublic
        ? 'selected-button'
        : 'animated-button'} nav-buttons"
      on:click={goToPublicApps}
    >
      Public Apps
    </button>
    {#if currentUserUID != ""}
      <button
        class="{selectedDev
          ? 'selected-button'
          : 'animated-button'} nav-buttons"
        on:click={goToDevDashboard}
      >
        Developer Dashboard
      </button>
    {/if}
  </div>
  <slot />
</div>

<style>
  .content-container {
    margin: 0 120px;
    color: #282936;
  }

  .nav-buttons-container {
    margin: 24px 0;
    display: flex;
    justify-content: center;
  }

  .animated-button {
    color: #282936;
    text-decoration: none;
    font-size: 26px;
    border: none;
    background: none;
    font-weight: 600;
    margin: 0 60px;
  }

  .animated-button::before {
    margin-left: auto;
  }

  .animated-button::after,
  .animated-button::before {
    content: "";
    width: 0%;
    height: 2px;
    background: #005be3;
    display: block;
    transition: 0.5s;
  }

  .animated-button:active {
    color: #003155;
    scale: 0.95;
  }

  .animated-button:hover::after,
  .animated-button:hover::before {
    width: 100%;
  }

  .selected-button {
    color: #005be3;
    text-decoration: none;
    font-size: 26px;
    border: none;
    background: none;
    font-weight: 600;
    margin: 0 60px;
  }

  .selected-button::after,
  .selected-button::before {
    content: "";
    height: 2px;
    background: #005be3;
    display: block;
    transition: 0.5s;
  }

  .selected-button:active {
    color: #003155;
    scale: 0.95;
  }
</style>
