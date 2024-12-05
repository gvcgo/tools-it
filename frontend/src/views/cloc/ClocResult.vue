<script setup lang="ts">
import { useRoute } from 'vue-router';

const route = useRoute();
const clocResult = {
    result: JSON.parse(route.query?.result?.toString() || "")
}
</script>

<template>
    <div v-if="clocResult.result.languages">
        <el-table border show-summary :data="clocResult.result.languages" style="width: 100%" :default-sort="{ prop: 'name', order: 'descending' }" >
            <el-table-column property="name" label="Name" width="240" sortable show-overflow-tooltip />
            <el-table-column property="files" label="Files" width="120" sortable  />
            <el-table-column property="code" label="Code" width="120" sortable  />
            
            <el-table-column property="comment" label="Comment" width="120" />
            <el-table-column property="blank" label="Blank" width="120" />
        </el-table>
    </div>
    <div v-else>
        <el-table border show-summary :data="clocResult.result.files" style="width: 100%" :default-sort="{ prop: 'language', order: 'ascending' }">
            <el-table-column property="name" label="Name" width="240" sortable show-overflow-tooltip />
            <el-table-column property="language" label="Language" width="120" sortable />
            <el-table-column property="code" label="Code" width="120" sortable />
            
            <el-table-column property="comment" label="Comment" width="120" />
            <el-table-column property="blank" label="Blank" width="120" />
        </el-table>
    </div>
  
</template>

<style scoped>

</style>