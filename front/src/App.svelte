<script>
  /*fetch("/meta/services")
    .then(res => res.json())
    .then(result => {
      console.log(result);
    })
    .catch(err => {
      console.log(`ERROR: ${err}`);
    });*/

  let username = "Elham Aryanpur";
  let theme = localStorage.getItem("theme");
  if (theme === null) {
    theme = "blue";
  }
</script>

<style>
  .table {
    width: 100%;
    height: 100vh;
    text-align: center;
    border-spacing: 20px 20px;
  }

  .menu {
    width: 15vw;
  }

  .menu-list {
    height: 100%;
    width: 100%;
    padding-top: 15px;
    padding-bottom: 15px;
    overflow: auto;
  }

  .menu-main {
    height: 100%;
    width: 100%;
  }
</style>

<svelte:head>
  <link rel="stylesheet" href="/themes/{theme}.css" />
</svelte:head>

<table class="table unselectable">
  <tr>
    <td class="card menu">
      <table class="menu-main">
        <tr>
          <div class="menu-list">
            <button class="box">Test1</button>
            <button class="box">Test1</button>
            <button class="box">Test1</button>
            <button class="box">Test1</button>
            <button class="box">Test1</button>
            <button class="box">Test1</button>
            <button class="box">Test1</button>
            <button class="box">Test1</button>
            <button class="box">Test1</button>
            <button class="box">Test1</button>
            <button class="box">Test1</button>
            <button class="box">Test1</button>
            <button class="box">Test1</button>
          </div>
        </tr>
      </table>
    </td>
    <td class="card">REEE</td>
  </tr>
</table>
