<script setup lang="ts">
import { reactive } from 'vue'
import { useRouter } from 'vue-router';
import { SelectDirectory, Cloc } from "../../../wailsjs/go/cloc/AppCloc";
import { Search} from '@element-plus/icons-vue'

// do not use same name with ref
const clocForm = reactive({
  byFile: false,
  skipDuplicated: true,
  sortTag: "code",
  excludeExt: "",
  includeLang: "",
  match: "",
  notMatch: "",
  matchDir: "",
  notMatchDir: "",
  paths: "",
})


const router = useRouter()

const onSubmit = () => {
    Cloc(
        clocForm.byFile,
        clocForm.skipDuplicated,
        clocForm.sortTag,
        clocForm.excludeExt,
        clocForm.includeLang,
        clocForm.match,
        clocForm.notMatch,
        clocForm.matchDir,
        clocForm.notMatchDir,
        clocForm.paths
    ).then((res) => {
        router.push({
            path: '/clocResult',
            query: {
                result: res,
            }
        })
    })
}

const onCancel = () => {
    clocForm.paths = ""
    clocForm.byFile = false
    clocForm.skipDuplicated = true
    clocForm.sortTag = "code"
    clocForm.excludeExt = ""
    clocForm.includeLang = ""
    clocForm.match = ""
    clocForm.notMatch = ""
    clocForm.matchDir = ""
    clocForm.notMatchDir = ""
}

const selectDir = () => {
    SelectDirectory().then((res) => {
        clocForm.paths = res
    })
}

</script>

<template>
    <el-form :model="clocForm" label-width="auto" style="max-width: 600px" class="item-center">
        <el-form-item label="Project Path">
            <el-input  v-model="clocForm.paths" >
                <template #append>
                    <el-button @click="selectDir" :icon="Search"></el-button>
                </template>
            </el-input>
        </el-form-item>
        <el-form-item label="Count By File">
            <el-switch v-model="clocForm.byFile" />
        </el-form-item>
        <el-form-item label="Skip Duplicated">
            <el-switch v-model="clocForm.skipDuplicated" />
        </el-form-item>
        <el-form-item label="Sort Tag">
            <el-select v-model="clocForm.sortTag" placeholder="please select your sort tag.">
                <el-option label="By Code" value="code" />
                <el-option label="By Name" value="name" />
                <el-option label="By Files" value="files" />
                <el-option label="By Blank" value="blank" />
                <el-option label="By Comment" value="comment" />
            </el-select>
        </el-form-item>
        <el-form-item label="Exclude Extensions">
            <el-input placeholder=".java,.go,.py" v-model="clocForm.excludeExt" />
        </el-form-item>
        <el-form-item label="Include Languages">
            <el-input placeholder="java,go,python" v-model="clocForm.includeLang" />
        </el-form-item>
        <el-form-item label="File Name to Match">
            <el-input placeholder="your regular expression" v-model="clocForm.match" />
        </el-form-item>
        <el-form-item label="File Name Not to Match">
            <el-input placeholder="your regular expression" v-model="clocForm.notMatch" />
        </el-form-item>
        <el-form-item label="Dir Name to Match">
            <el-input placeholder="your regular expression" v-model="clocForm.matchDir" />
        </el-form-item>
        <el-form-item label="Dir Name Not to Match">
            <el-input placeholder="your regular expression" v-model="clocForm.notMatchDir" />
        </el-form-item>

        <el-form-item>
            <el-button type="primary" @click="onSubmit">Count</el-button>
            <el-button @click="onCancel">Cancel</el-button>
        </el-form-item>
    </el-form>
</template>
  
<style scoped>
/* ::v-deep是vue3提供的深度选择器，.el-form-item__label是element-plus官方定义的类 */
::v-deep .el-form-item__label{
  color: white;
  font-size: 18px;
}
</style>
