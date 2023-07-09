<template>
  <div class="common-layout">
    <el-header class="admin-header">Header</el-header>
    <el-container>
      <el-aside width="240px" class="admin-aside" id="admin-aside">
        <div>
          <el-scrollbar>
            <el-menu>
              <el-sub-menu index="1">
                <template #title>
                  <el-icon>
                    <message/>
                  </el-icon>
                  Navigator One
                </template>
                <el-menu-item-group>
                  <template #title>Group 1</template>
                  <el-menu-item index="1-1">Option 1</el-menu-item>
                  <el-menu-item index="1-2">Option 2</el-menu-item>
                </el-menu-item-group>
                <el-menu-item-group title="Group 2">
                  <el-menu-item index="1-3">Option 3</el-menu-item>
                </el-menu-item-group>
                <el-sub-menu index="1-4">
                  <template #title>Option4</template>
                  <el-menu-item index="1-4-1">Option 4-1</el-menu-item>
                </el-sub-menu>
              </el-sub-menu>
              <el-sub-menu index="2">
                <template #title>
                  <el-icon>
                    <icon-menu/>
                  </el-icon>
                  Navigator Two
                </template>
                <el-menu-item-group>
                  <template #title>Group 1</template>
                  <el-menu-item index="2-1">Option 1</el-menu-item>
                  <el-menu-item index="2-2">Option 2</el-menu-item>
                </el-menu-item-group>
                <el-menu-item-group title="Group 2">
                  <el-menu-item index="2-3">Option 3</el-menu-item>
                </el-menu-item-group>
                <el-sub-menu index="2-4">
                  <template #title>Option 4</template>
                  <el-menu-item index="2-4-1">Option 4-1</el-menu-item>
                </el-sub-menu>
              </el-sub-menu>
              <el-sub-menu index="3">
                <template #title>
                  <el-icon>
                    <setting/>
                  </el-icon>
                  Navigator Three
                </template>
                <el-menu-item-group>
                  <template #title>Group 1</template>
                  <el-menu-item index="3-1">Option 1</el-menu-item>
                  <el-menu-item index="3-2">Option 2</el-menu-item>
                </el-menu-item-group>
                <el-menu-item-group title="Group 2">
                  <el-menu-item index="3-3">Option 3</el-menu-item>
                </el-menu-item-group>
                <el-sub-menu index="3-4">
                  <template #title>Option 4</template>
                  <el-menu-item index="3-4-1">Option 4-1</el-menu-item>
                </el-sub-menu>
              </el-sub-menu>
            </el-menu>
          </el-scrollbar>
        </div>
      </el-aside>
      <el-container>
        <el-main>Main</el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script setup>
import {Menu as IconMenu, Message, Setting} from '@element-plus/icons-vue'

</script>

<style scoped>
.admin-aside {
  height: 100vh;
  border-right: solid 1px lavender;
}

.admin-header {
  height: 60px;
  background-color: deeppink;
}

.admin-logo-container {
  height: 60px;
  background-color: darkturquoise;
}


</style>
