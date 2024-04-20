<script>
  import { onMount } from "svelte";
  import AppCard from "../../lib/AppCard.svelte";
  import { currentUser } from "../../stores.js";

  let currentUserUID = "";
  currentUser.subscribe((value) => {
    currentUserUID = value;
  });

  let appName = "";
  let baseURL = "";

  let apps = [];

  function getDevApps() {
    fetch("http://127.0.0.1:3000/getDevApps", {
      method: "POST",
      headers: {
        "Content-Type": "text/plain",
      },
      body: currentUserUID,
    })
      .then((response) => {
        return response.json();
      })
      .then((data) => {
        apps = data;
      })
      .catch((error) => {
        console.error("There was a problem with the fetch operation:", error);
      });
  }

  function addApp() {
    const newApp = {
      name: appName,
      developer: currentUserUID,
      baseURL: baseURL,
      endpoints: [],
      status: "Down",
    };

    addAppToDB(newApp);
    apps = [...apps, newApp];
  }

  function addAppToDB(app) {
    fetch("http://127.0.0.1:3000/addApp", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(app),
    })
      .then((response) => {
        console.log(response);
        response.json();
      })
      .catch((error) => {
        console.error("There was a problem with the fetch operation:", error);
      });
  }

  $: {
    console.log(currentUserUID);
    getDevApps();
  }
</script>

{#if currentUserUID == ""}
  <div style="display: flex; justify-content:center; margin-top: 120px">
    <h2>You cannot view this page if you are not logged in.</h2>
  </div>
{:else}
  <div class="header">
    <h1>Developer Dashboard</h1>
    <div class="btn-group dropend">
      <button
        type="button"
        class="btn btn-success dropdown-toggle"
        data-bs-toggle="dropdown"
        aria-expanded="false"
        style="height: 40px; text-align: center; display: flex; align-items: center; justify-content: center;"
      >
        Add an app
      </button>
      <form
        class="dropdown-menu p-3"
        style="width: 300px"
        on:submit|preventDefault={addApp}
      >
        <div class="mb-3">
          <label for="appName" class="form-label">App name:</label>
          <input
            type="text"
            class="form-control"
            id="appName"
            placeholder="Example app"
            bind:value={appName}
          />
        </div>
        <div class="mb-3">
          <label for="baseurl" class="form-label">BaseURL:</label>
          <input
            type="text"
            class="form-control"
            id="baseurl"
            placeholder="https://example.com/"
            bind:value={baseURL}
          />
        </div>
        <button type="submit" class="btn btn-primary">Add app</button>
      </form>
    </div>
  </div>

  {#if apps.length > 0}
    <div class="list">
      {#each apps as app}
        <AppCard
          appName={app.name}
          baseURL={app.baseURL}
          endpoints={app.endpoints}
          status={app.status}
        />
      {/each}
    </div>
  {/if}
{/if}

<style>
  .header h1 {
    margin-right: 32px;
  }
</style>
