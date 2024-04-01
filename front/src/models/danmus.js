// src/models/danmus.js
export default {
    namespace: 'danmus',
    state: {
      list: [], // 存储弹幕列表
    },
    reducers: {
      // 添加新的弹幕到列表
      addDanmu(state, { payload }) {
        return {
          ...state,
          list: [...state.list, payload],
        };
      },
    },
    effects: {
      // 可以添加异步获取弹幕数据的逻辑
    },
  };
  