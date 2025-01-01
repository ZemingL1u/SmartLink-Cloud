// event-bus.js
import { reactive, readonly } from 'vue';

// 创建一个响应式状态对象
const state = reactive({
    open: false,
});

export default {
    // 提供一个方法来更改 open 状态
    // changeOpen(value) {
    //     state.open = value;
    //     console.log('open state changed to:', state.open);
    // },
    changeOpen(value) {
        // 使用 Vue 的 nextTick 等待更新
        import('vue').then(vue => {
            vue.nextTick(() => {
                state.open = value;
                console.log('open state changed to:', state.open);
            });
        });
    },
    // 提供当前状态的只读版本
    state: readonly(state),
};