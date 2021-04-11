<template>
  <div id="app">
    <header>
      <nav>
        <a href="/"><img alt="Logo" src="./assets/logo.png" height="70" /></a>
        PCS-CLI WEB UI
        <ul>
          <!-- HERE WE LOOP ALL TAGS FROM THE COMMANDS -->
          <li>Platform</li>
          <li><a href="#">Service/ API</a></li>
          <li>
            <a href="#">Provider</a>
            <ul>
              <li><a href="#">Sublink with a long name</a></li>
              <li><a href="#">Short sublink</a></li>
            </ul>
          </li>
        </ul>
      </nav>
      <h1>PCS-CLI WEB UI</h1>
      <p>A simple command building tool for the PCS-CLI tool!</p>
      <br />
      <p>
        <a href="#"><i>PCS-CLI Documentation</i></a
        ><a href="#"><b>Commands &rarr;</b></a>
      </p>
    </header>
    <main>
      <CommandHelper />
      <hr />
    </main>
    <footer>
      <hr />
      <p>
        <small>&copy; 2021 Nike, Inc. All Rights Reserved</small>
      </p>
    </footer>
  </div>
</template>

<script>
import CommandHelper from "./components/CommandHelper.vue";

export default {
  name: "App",
  components: {
    CommandHelper,
  },
};
</script>

<style scoped>
nav a {
  padding: 3px;
}
nav a:visited {
  color: #fe4712;
}
nav a:hover {
  color: #ffffff;
  background: #fe4712;
  text-decoration: none;
}
nav a:hover img {
  background: transparent;
  text-decoration: none;
}
pre code {
  background: #212f3d;
  color: #ffffff;
  font-family: console;
}
</style>
