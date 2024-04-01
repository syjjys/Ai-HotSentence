import dva from 'dva';
import './index.css';
import barrages from './models/barrages';

// 1. 初始化
const app = dva();

// 2. Model
const context = require.context('./models', false, /\.js$/);

const autoLoadModels = app1 => {
  context.keys().forEach(key => {
    const model = context(key).default;
    app1.model(model);
  }); 
}

app.model(barrages)

// 3. Router
app.router(require('../src/router').default);

// 4. 启动应用
app.start('#root');
