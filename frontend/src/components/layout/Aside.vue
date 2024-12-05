<script setup lang="ts">
import { ref } from  "vue";

</script>

<template>
    <el-aside width="200px" >
        <el-menu  router>
            <el-menu-item  index="/clocQuery"><el-icon><IEpDataAnalysis /></el-icon><span>cloc</span></el-menu-item>
        </el-menu>
    </el-aside>
</template>
  
<style scoped>
.el-side {
    background-color: #D4D7DE;
}

.el-menu {
    background-color: #D4D7DE;
    .el-menu-item {
        background-color: #D4D7DE;
    }
}
</style>