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
    <form class="dropdown-menu p-3" style="width: 300px">
      <div class="mb-3">
        <label for="appName" class="form-label">Password</label>
        <input
          type="text"
          class="form-control"
          id="appName"
          placeholder="Example app"
        />
      </div>
      <div class="mb-3">
        <label for="baseurl" class="form-label">BaseURL:</label>
        <input
          type="text"
          class="form-control"
          id="baseurl"
          placeholder="https://example.com/"
        />
      </div>
      <button type="submit" class="btn btn-primary">Add app</button>
    </form>
  </div>
</div>

<div class="list">
  <div class="card border-secondary" style="width: 20rem;">
    <div class="card-body">
      <h2 class="card-title" style="margin-bottom: 12px">App name</h2>
      <hr />
      <p class="card-text">https://example.com/</p>
      <p class="card-text">Endpoints: 2</p>
      <p class="card-text">Status: Up</p>
      <a href="#" class="btn btn-primary">View details</a>
    </div>
  </div>
</div>

<style>
  .header {
    display: flex;
    flex-direction: row;
    align-items: center;
    margin-bottom: 46px;
  }
  .header h1 {
    margin-right: 32px;
  }
  .list {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(20rem, 1fr));
  }
  .card {
    margin-bottom: 32px;
  }
</style>
