<template>
  <el-row :gutter="0">
    <el-col :md=2 class="app-area">
      <img src="@/assets/logo.png" alt="logo" class="logo">
    </el-col>
    <el-col :md="16" class="app-area">
      <div>
        <el-link type="primary" :underline="false" :icon="House"><router-link to="/">Home</router-link></el-link>
        <el-link type="primary" :underline="false" :icon="Compass"><router-link to="/baseline">Baseline</router-link></el-link>
        <el-link type="primary" :underline="false" :icon="Notebook"><router-link to="/blogs">Blogs</router-link></el-link>
        <el-link type="primary" :underline="false" :icon="Pointer"><router-link to="/apply">Apply</router-link></el-link>
      </div>
      <el-dropdown>
        <el-link type="primary" :underline="false" style="margin-top: 5px;" :icon="Service"><router-link to="">Administrator</router-link></el-link>
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item><el-icon><Pointer /></el-icon><router-link to="/admin/applies">Applies Manager</router-link></el-dropdown-item>
            <el-dropdown-item><el-icon><User /></el-icon><router-link to="/admin/users">Users Manager</router-link></el-dropdown-item>
            <el-dropdown-item><el-icon><Operation /></el-icon><router-link to="/admin/projects">Projects Manager</router-link></el-dropdown-item>
            <el-dropdown-item><el-icon><Tickets /></el-icon><router-link to="/admin/flows">Flows Manager</router-link></el-dropdown-item>
            <el-dropdown-item divided><el-icon><MagicStick /></el-icon><router-link to="/admin/spiders">Spider Manager</router-link></el-dropdown-item>
            <el-dropdown-item><el-icon><Calendar /></el-icon><router-link to="/admin/schedulers">Schedule Manager</router-link></el-dropdown-item>
            <el-dropdown-item divided><el-icon><Setting /></el-icon><router-link to="/admin/web/settings">Web Settings</router-link></el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
      <div>
        <input v-model="input1" placeholder="Search OSS" class="search-input"/>
      </div>
    </el-col>
    <el-col :md="6" class="user-area">
      <div class="notification-area">
        <el-badge :value="44" class="item" type="danger">
          <el-icon><router-link to="/notifications"><Bell /></router-link></el-icon>
        </el-badge>
        <el-badge :value="3" class="item" type="warning">
          <el-icon><router-link to="/reviews"><Guide /></router-link></el-icon>
        </el-badge>
      </div>
      <div>
        <el-dropdown>
          <el-avatar shape="circle" :size="30" :src="avatarUrl" />
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item>SignIn As: wuy37</el-dropdown-item>
              <el-dropdown-item divided style="align-items: center;justify-content: center;"><el-button type="plain" :icon="Headset">Set Status</el-button></el-dropdown-item>
              <el-dropdown-item divided><el-icon><Pointer /></el-icon><router-link to="/user/applies">Your Apply</router-link></el-dropdown-item>
              <el-dropdown-item><el-icon><Tickets /></el-icon><router-link to="/user/flows">Your Flows</router-link></el-dropdown-item>
              <el-dropdown-item><el-icon><User /></el-icon><router-link to="/user/profile">Your Profile</router-link></el-dropdown-item>
              <el-dropdown-item><el-icon><Notebook /></el-icon><router-link to="/user/blogs">Your Blogs</router-link></el-dropdown-item>
              <el-dropdown-item divided><el-icon><SwitchButton /></el-icon><router-link to="/auth/login">Login System</router-link></el-dropdown-item>
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
import { Bell, Guide, House, Compass, Search, Notebook, Headset, Pointer, Setting, Service, SwitchButton, Operation, User, Calendar, Tickets, MagicStick} from '@element-plus/icons-vue'

const state = reactive({
  avatarUrl:
      'https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png',
})
const { avatarUrl } = toRefs(state)
const input1 = ref('')

</script>

<style scoped>
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
  color: rgb(64, 158, 255);
}

a {
  text-decoration: none;
  color: rgb(64, 158, 255);
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
