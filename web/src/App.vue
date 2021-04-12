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
      <div v-for="(cmd, index) in commandsList.commands" :key="index">
        <CommandHelper :command="cmd" />
      </div>
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
  methods: {
    copyEventHandler() {},
    copyToClipboard(text) {
    if (window.clipboardData && window.clipboardData.setData) {
        // Internet Explorer-specific code path to prevent textarea being shown while dialog is visible.
        return window.clipboardData.setData("Text", text);

    }
    else if (document.queryCommandSupported && document.queryCommandSupported("copy")) {
        var textarea = document.createElement("textarea");
        textarea.textContent = text;
        textarea.style.position = "fixed";  // Prevent scrolling to bottom of page in Microsoft Edge.
        document.body.appendChild(textarea);
        textarea.select();
        try {
            return document.execCommand("copy");  // Security exception may be thrown by some browsers.
        }
        catch (ex) {
            console.warn("Copy to clipboard failed.", ex);
            return false;
        }
        finally {
            document.body.removeChild(textarea);
        }
    }
}
  }
  data() {
    return {
      commandsList: {
        commands: [
          {
            name: "register-platform",
            description: "Register a platform with PCS",
            flags: [
              {
                name: "platform-name",
                description: "Name of the platform",
                value: "avalue",
              },
              {
                name: "version",
                description: "Attribute value of the platform version",
                value: "",
              },
              {
                name: "capabilities",
                description: "List of capabilities which the platform has",
                value: "",
              }
            ]
          },
          {
            name: "register-provider",
            description: "Register and establish trust between PCS and a platform provider",
            flags: [
              {
                name: "platform-name",
                description: "Name of the platform",
                value: "",
              },
              {
                name: "provider",
                description: "Name of the Provider. For example 'AWS'",
                value: "",
              },
              {
                name: "account-id",
                description: "Provider account ID",
                value: "",
              }
            ]
          },
          {
            name: "register-api",
            description: "Several steps taken here to establish the important pieces of an API: Ensure IGW exists, creates ALB, ....",
            flags: [
              {
                name: "platform-name",
                description: "Name of the platform",
                value: "avalue",
              },
              {
                name: "api-name",
                description: "Name of the API service you are creating",
                value: "",
              },
              {
                name: "version",
                description: "The version of the API Service",
                value: "",
              },
              {
                name: "vpc-id",
                description: "VPC ID",
                value: "",
              },
              {
                name: "target-group-arn",
                description: "Target Group ARN",
                value: "",
              },
              {
                name: "hosted-zone-id",
                description: "hosted zone id of the service",
                value: "",
              }
            ]
          }
        ]
      }
    };
  }
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
