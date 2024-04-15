<script>
  export let color = "yellow";
</script>

<div
  class="status-box {color === 'red'
    ? 'red'
    : color === 'green'
      ? 'green'
      : 'yellow'}"
></div>

<style>
  .status-box {
    border-radius: 4px;
    width: 20px;
    height: 20px;
    margin-bottom: 17px;
    margin-right: 4px;
  }
  .green {
    background-color: #58ba16;
  }
  .yellow {
    background-color: #ffbd44;
  }
  .red {
    background-color: #f04532;
  }
</style>
