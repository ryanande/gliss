<template>
  <div>
    <details>
      <summary>{{ command.name }}</summary>
      <p>{{ command.description }}</p>
      <form>
        <header>
          <h2>{{ command.name }} Flags</h2>
        </header>
        <div v-for="(flag, index) in command.flags" :key="index">
          <label for="input1">{{ index }} : {{ flag.name }}</label>
          <input type="text" v-model="flag.value" />
        </div>
        <pre>
          <code>./pcs {{ command.name }} {{ getcommandFlagValues }}
          </code>
        </pre>
      </form>
    </details>
    <br />
  </div>
</template>

<script>
export default {
  name: "CommandHelper",
  props: {
    command: {
      type: Object,
      default: () => ({
        name: "< comandname >",
        description: "< detailed flag description here... >",
        flags: [
          { name: "platform-name-a", description: "< flag description here... >", value: "avalue" },
          { name: "platform-name-b", description: "< flag description here... >", value: "bvalue" }
        ]
      })
    },
  },
  computed: {
    getcommandFlagValues() {
      return this.command.flags.map(function(f){
          return "--" + f.name + " " + f.value;
      }).join(" ");
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
