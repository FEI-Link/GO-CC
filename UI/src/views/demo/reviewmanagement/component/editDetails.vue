
<template>
	<div class="system-edit-post-container">
		<el-dialog :title="(formData.id !== 0 ? '修改' : '添加') + '评测内容'" v-model="isShowDialog" width="1000px">
			<el-form ref="formRef" :model="formData" :rules="rules" size="default" label-width="180px">
				<!-- <el-form-item v-if="formData.id === 0"  label="用户ID" prop="user_id">
					<el-input v-model="formData.user_id" placeholder="请输入唯一用户ID" />
				</el-form-item> -->
				<el-form-item label="评测内容" prop="neirong">
					<el-input v-model="formData.neirong" placeholder="请输入评测内容" />
				</el-form-item>
				<!-- <el-form-item label="详细内容" prop="study_de">
					<el-input v-model="formData.study_de" placeholder="请输入学历详细信息" />
				</el-form-item> -->
	
			</el-form>
			<template #footer>
				<span class="dialog-footer">
					<el-button @click="onCancel" size="default">取 消</el-button>
					<el-button type="primary" @click="onSubmit" size="default" :loading="loading">{{ formData.id !== 0 ?
			'修 改' : '添 加' }}
					</el-button>
				</span>
			</template>
		</el-dialog>
	</div>
</template>


<script setup lang="ts">
import { reactive, toRefs, ref, unref } from 'vue';
import {editpingce } from '/@/api/demo/pingce';
import { ElMessage } from "element-plus";

interface FormState {
	id: number;
	study: string;
	neirong: string;
	user_id: string;
}
interface ZhiShiKuState {
	isShowDialog: boolean;
	loading: boolean;
	formData: FormState;
	rules: {}
}

const formRef = ref<HTMLElement | null>(null);
const state = reactive<ZhiShiKuState>({
	loading: false,
	isShowDialog: false,
	formData: {
		id: 0,
		study:'',
		user_id: '',
		neirong: '',
	},
	// 表单校验
	rules: {
		userName: [
			{ required: true, message: "姓名不能为空", trigger: "blur" }
		],
		position: [
			{ required: true, message: "职务不能为空", trigger: "blur" }
		],
	}
})

const resetForm = () => {
	state.formData = {
		id: 0,
		study:'',
		user_id: '',
		neirong: '',
	}
};

const emit = defineEmits<{
	(e: "dataList"): void;
}>();

const openDialog = (row: FormState | null) => {
	resetForm();
	if (row) {
		state.formData = row;
	}
	state.isShowDialog = true;
};

// 关闭弹窗
const closeDialog = () => {
	state.isShowDialog = false;
};
// 取消
const onCancel = () => {
	closeDialog();
};

// 新增
const onSubmit = () => {
	const formWrap = unref(formRef) as any;
	if (!formWrap) return;
	formWrap.validate((valid: boolean) => {
		if (valid) {
			state.loading = true;
			// 发送到后台
			if (state.formData.id === 0) {
				// 发送到后台
				console.log("测试是否调用了添加函数");
				addstudy(state.formData).then(() => {
					ElMessage.success('新职务信息添加成功');
					closeDialog(); // 关闭弹窗
					emit('dataList');// 给父窗口发送事件
				}).finally(() => {
					state.loading = false;
				})
			} else {
				// 发送到后台
				editpingce(state.formData).then(() => {
					ElMessage.success('修改成功');
					closeDialog(); // 关闭弹窗
					emit('dataList');// 给父窗口发送事件
				}).finally(() => {
					state.loading = false;
				})
			}
		}
	});
};

let { loading, isShowDialog, formData, rules } = toRefs(state);

defineExpose({
	openDialog,
})
</script>

<style></style>