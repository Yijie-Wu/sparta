<template>
  <el-row class="login-container">
    <el-col :xs="8" :sm="8" :md="8" :lg="8" :xl="8">
      <div class="login-box">1</div>
    </el-col>
    <el-col :xs="16" :sm="16" :md="16" :lg="16" :xl="16">
      <div class="login-image-box">2</div>
    </el-col>
  </el-row>
</template>

<script>

</script>

<style scoped>
.login-container {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

.login-box {
  flex: 1;
  background-color: paleturquoise;
}

.login-image-box {
  flex: 1;
  background-color: deeppink;
}
</style>
