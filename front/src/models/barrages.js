export default {
    namespace: 'barrages',
    state: {
      messages: [],
      data:[]
    },
    reducers: {
      saveMessages(state, { payload: messages1 }) {
        return { ...state, messages :[...messages1] };
      },
      saveDatas(state, { payload: data1 }) {
        return { ...state, data :[...data1] };
      },
      addLike(state, {payload}) {
        const updated = state.messages.map(mes => mes.sayId === payload.sayId ? {...mes, likeNum: (parseInt(mes.likeNum, 10) + 1).toString()} : mes)
        return {...state, messages:[...updated]}
      },
      addLikeData(state, {payload}) {
        const updated = state.data.map(mes => mes.sayId === payload.sayId ? {...mes, likeNum: (parseInt(mes.likeNum, 10) + 1).toString()} : mes)
        return {...state, data:[...updated]}
      },
      addComment(state, {payload}) {
        const updated = state.messages.map(mes => mes.sayId === payload.sayId ? {...mes, commentNum: (parseInt(mes.commentNum, 10) + 1).toString()} : mes)
        return {...state, messages:[...updated]}
      }
    },
    effects: {
    //   *fetchMessages({ payload }, { put }) {
    //     // 这里可以发起请求获取弹幕数据
    //     // 假设获取到的数据为 messages
    //     yield put({ type: 'saveMessages', payload: messages });
    //   },
    },
  };
  