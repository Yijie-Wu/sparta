<template>
  <el-row :gutter="0">
    <el-col :xs="8" :sm="6" :md="4" :lg="3" :xl="1">
      <div>
          <router-link to="/">Home</router-link> |
          <router-link to="/about">About</router-link>
      </div>
    </el-col>
    <el-col :xs="4" :sm="6" :md="8" :lg="9" :xl="11">
      <div>q</div>
    </el-col>
    <el-col :xs="4" :sm="6" :md="8" :lg="9" :xl="11">
      <div>q</div>
    </el-col>
    <el-col :xs="8" :sm="6" :md="4" :lg="3" :xl="1">
      <div>q</div>
    </el-col>
  </el-row>
  <router-view/>
</template>

<script setup>

</script>

<style scoped>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

.el-row {
  background-color: seagreen;
}

.el-col {
  height: 60px;
}

</style>
