<script>
  import { goto } from "$app/navigation";
  export let appName = "Example app";
  export let baseURL = "https://example.com/";
  export let endpoints = 0;
  export let status = "Down";

  function openAppPage() {
    goto("../appPage");
  }
</script>

<div class="list">
  <div class="card border-secondary" style="width: 20rem;">
    <div class="card-body">
      <h2 class="card-title" style="margin-bottom: 12px">{appName}</h2>
      <hr />
      <p class="card-text">{baseURL}</p>
      <p class="card-text">Endpoints: {endpoints}</p>
      <p class="card-text">Status: {status}</p>
      <button class="btn btn-primary" on:click={openAppPage}
        >View details</button
      >
    </div>
  </div>
</div>

<style>
  .card {
    margin-bottom: 32px;
  }
</style>
