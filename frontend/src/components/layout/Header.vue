<script setup lang="ts">
import { Quit, WindowHide, WindowMinimise, WindowToggleMaximise } from "../../../wailsjs/runtime";
import { ref } from  "vue";
import { GetAppName } from "../../../wailsjs/go/cloc/AppCloc";

let appName =  ref("");
GetAppName().then((name) =>{
  appName.value = name;
})
</script>

<template>
    <el-header class="flex px-1">
      <div class="flex items-center">{{appName}}</div>
      <div class="flex-grow h-full" style="--wails-draggable:drag"></div>
      <div class="space-x-2 flex">
        <span class="cursor-pointer flex items-center" @click="WindowHide()">
          <el-icon size="25px"><i-ep-ArrowDown /></el-icon>
        </span>
        <span class="cursor-pointer flex items-center" @click="WindowMinimise()">
          <el-icon size="25px"><i-ep-RemoveFilled /></el-icon>
        </span>
        <span class="cursor-pointer flex items-center" @click="WindowToggleMaximise()">
          <el-icon size="25px"><i-ep-Rank /></el-icon>
        </span>
        <span class="cursor-pointer flex items-center" @click="Quit()">
          <el-icon size="25px"><i-ep-CloseBold /></el-icon>
        </span> 
      </div>
    </el-header>
</template>
  
<style scoped>
.el-header {
  height: 35px;
}
</style>