// src/components/DanmuComponent.js
import React, { useEffect } from 'react';
import { connect } from 'dva';

const DanmuComponent = ({ dispatch, danmus }) => {
  useEffect(() => {
    dispatch({type: 'danmus/addDanmu', payload: {content:"666"}})
    // 可以在这里通过dispatch调用model的effect来异步获取弹幕数据
    // 例如: dispatch({ type: 'danmus/fetchDanmus' });
  }, [dispatch]);

  return (
    <div className="danmu-container">
      {danmus.list.map((danmu, index) => (
        <div key={index} className="danmu" style={{ /* 弹幕的CSS样式 */ }}>
          {danmu.content}
        </div>
      ))}
    </div>
  );
};

export default connect(({ danmus }) => ({
  danmus,
}))(DanmuComponent);
