export default function shuffleArray(array) {
  // 首先创建原数组的副本
  let arrayCopy = [...array];
  // 然后对副本进行随机排序
  for (let i = arrayCopy.length - 1; i > 0; i--) {
    // 生成从0到i的随机索引
    const j = Math.floor(Math.random() * (i + 1));
    // 交换元素arrayCopy[i]和arrayCopy[j]
    [arrayCopy[i], arrayCopy[j]] = [arrayCopy[j], arrayCopy[i]];
  }
  return arrayCopy;
}
