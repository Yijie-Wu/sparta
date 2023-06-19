// 从vue导入 createApp 函数
import { createApp } from 'vue'
// 导入 App 根组件
import App from './App.vue'
// 导入 router
import router from './router'
// 导入 store
import store from './store'
// 导入 ElementPlus
import ElementPlus from 'element-plus'
// 导入 element-plus css
import 'element-plus/dist/index.css'

// 创建 app 实例
const app = createApp(App)
// app 绑定各个组件并挂载到 #app 下面
app.use(store).use(router).use(ElementPlus).mount('#app')
