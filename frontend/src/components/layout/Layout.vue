<script setup lang="ts">
import Header from './Header.vue'
import Aside from './Aside.vue'
import Stepper from './Stepper.vue';
import { RouterView } from 'vue-router';
</script>

<template>
    <div class="common-layout">
        <el-container class="header-and-main">
            <!-- <el-header>Header</el-header> -->
            <Header class="flex h-[25px] w-full items-center bg-blue-500/50"></Header>
            <el-container class="aside-and-main">
                <Aside class="aside" />
                <el-container class="stepper-and-main">
                    <Stepper />
                    <el-main>
                        <el-scrollbar><RouterView /></el-scrollbar>
                    </el-main>
                </el-container>
            </el-container>
        </el-container>
    </div>
</template>
  
<style scoped>
.header-and-main {
    flex-direction: column;
}
.aside-and-main {
    flex-direction: row;
    height: calc(100vh - 35px);
}
.stepper-and-main {
    flex-direction: column;
}
.aside {
    background-color: #D4D7DE;
}
.el-main {
    margin: 0;
    padding: auto;
}

.el-scrollbar {
    margin: 0;
    padding: 0;
}
</style>