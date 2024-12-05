<script setup lang="ts">
import { useRouter } from 'vue-router'
 
const router = useRouter()
const goBack = () => {
    router.go(-1)
}
const goNext = () => {
    router.go(1)
}
</script>

<template>
    <el-header style="font-size: 25px" id="stepper" class="flex px-1">
        <el-icon  @click="goBack" class="cursor-pointer flex"><IEpCaretLeft /></el-icon> 
        <el-icon  @click="goNext" class="cursor-pointer flex left-2"><IEpCaretRight /></el-icon>
    </el-header>
</template>

<style scoped>
#stepper {
    height: 20px;
    margin: 0;
}
</style>