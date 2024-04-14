<script>
  import { goto } from "$app/navigation";

  let loginState = "LoginState";

  let selectedPublic = false;
  function goToPublicApps() {
    selectedDev = false;
    selectedPublic = !selectedPublic;
    goto("/");
  }

  let selectedDev = false;
  function goToDevDashboard() {
    selectedPublic = false;
    selectedDev = !selectedDev;
    goto("../devDashboard");
  }
</script>

<nav class="navbar navbar-expand-lg bg-primary" style="padding: 18px 32px">
  <div class="container-fluid">
    <h1 class="navbar-child" style="margin-bottom: 0">ItecMonitor</h1>
    <div class="dropdown">
      <button
        type="button"
        class="btn dropdown-toggle login-bttn"
        data-bs-toggle="dropdown"
        data-bs-display="static"
        aria-expanded="false"
        style="color: white"
      >
        loginState
      </button>
      <ul class="dropdown-menu dropdown-menu-lg-end">
        <li><button class="login-options">Login with Google</button></li>
        <li><button class="login-options">Login with Github</button></li>
      </ul>
    </div>
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
    {#if loginState !== "Login"}
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
    margin: 0 72px;
    color: #282936;
  }

  .nav-buttons-container {
    margin: 24px 0;
  }

  .animated-button {
    color: #282936;
    text-decoration: none;
    font-size: 25px;
    border: none;
    background: none;
    font-weight: 600;
    margin-right: 64px;
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

  /* .animated-button:focus {
    transition: 0.3s;
    color: #005be3;
  } */

  .selected-button {
    color: #005be3;
    text-decoration: none;
    font-size: 25px;
    border: none;
    background: none;
    font-weight: 600;
    margin-right: 64px;
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
