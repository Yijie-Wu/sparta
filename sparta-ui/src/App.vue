<template>
  <el-row :gutter="0">
    <el-col :md=2 class="app-area">
      <img src="@/assets/logo.png" alt="logo" class="logo">
    </el-col>
    <el-col :md="16" class="app-area">
      <div>
        <el-link type="primary" :underline="false"><router-link to="/">Home</router-link></el-link>
        <el-link type="primary" :underline="false"><router-link to="/baseline">Baseline</router-link></el-link>
        <el-link type="primary" :underline="false"><router-link to="/blogs">Blogs</router-link></el-link>
        <el-link type="primary" :underline="false"><router-link to="/apply">Apply</router-link></el-link>
        <el-link type="primary" :underline="false"><router-link to="/admin">Administrator</router-link></el-link>
      </div>
      <div>
        <input v-model="input1" placeholder="Search OSS" class="search-input"/>
      </div>
    </el-col>
    <el-col :md="6" class="user-area">
      <div class="notification-area">
        <el-badge :value="44" class="item" type="danger">
          <el-icon><Bell /></el-icon>
        </el-badge>
        <el-badge :value="3" class="item" type="warning">
          <el-icon><Guide /></el-icon>
        </el-badge>
      </div>
      <div>
        <el-dropdown>
            <el-avatar shape="circle" :size="30" :src="avatarUrl" />
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item>Your Apply</el-dropdown-item>
              <el-dropdown-item>Your Flows</el-dropdown-item>
              <el-dropdown-item>Action 3</el-dropdown-item>
              <el-dropdown-item>Action 4</el-dropdown-item>
              <el-dropdown-item divided>Login</el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </div>
    </el-col>
  </el-row>
  <router-view/>
</template>

<script setup>
import { reactive, toRefs, ref } from 'vue'
import { Bell, Guide } from '@element-plus/icons-vue'

const state = reactive({
  avatarUrl:
      'https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png',
})
const { avatarUrl } = toRefs(state)
const input1 = ref('')

</script>

<style scoped>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

.el-row {
  background-color: rgba(217, 175, 238, 0.11);
}

.app-area {
  height: 60px;
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: left;
}

.user-area {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: right;
  padding-right: 15px;
}


.logo {
  width: 50px;
  height: 50px;
  margin: 5px 10px;
}

.el-dropdown {
  border: none;
  outline: none;
  border-radius: 50%;
}

.el-dropdown:hover {
  border: none;
  outline: none;
  border-radius: 50%;
}

.el-link {
  margin-right: 15px;
}

.el-link a {
  text-decoration: none;
  color: #03f2f8;
}

.search-input {
  height: 32px;
  min-width: 360px;
  border-radius: 16px;
  outline: none;
  border: solid 0.5px darkseagreen;
  padding: 0 15px;
}

.el-badge {
  margin-right: 30px;
}

.notification-area {
  margin-right: 1px;
}

</style>
