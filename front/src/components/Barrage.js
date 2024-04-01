import React, { useEffect, useRef, useState } from "react";
import "./Barrage.css"; // 引入弹幕样式
import { connect } from "dva";
import shuffleArray from "../utils/commonUtil";
import { Triangle } from "react-loader-spinner";
import {
  EditOutlined,
  EllipsisOutlined,
  SettingOutlined,
  HeartTwoTone,
  HeartFilled,
  HeartOutlined,
  LikeOutlined,
  MessageOutlined,
  RedoOutlined,
} from "@ant-design/icons";
import {
  Avatar,
  Card,
  Button,
  Flex,
  List,
  Divider,
  Skeleton,
  Switch,
  Drawer,
  Radio,
  Input,
  Modal,
  Space,
  Affix,
} from "antd";
import InfiniteScroll from "react-infinite-scroll-component";

const { Meta } = Card;
const { TextArea } = Input;
const Barrage = ({ dispatch, barrages }) => {
  const [open, setOpen] = useState(false);
  const [placement, setPlacement] = useState("left");
  const showDrawer = () => {
    setOpen(true);
  };
  const onChange = (e) => {
    setPlacement(e.target.value);
  };
  const onClose = () => {
    setOpen(false);
  };
  const [openModal, setOpenModal] = useState(false);

  const showModal = () => {
    setOpenModal(true);
  };

  const hideModal = () => {
    setOpenModal(false);
  };

  const [com, setComment] = useState("");
  const [nickNam, setNickName] = useState("孤独的狼");
  const [sayI, setSayId] = useState("");
  const [coms, setComments] = useState([""]);

  const Load = () => {
    return (
      <div
        style={{
          display: "flex",
          flexDirection: "column",
          justifyContent: "center",
          alignItems: "center",
          height: "100vh", // 这里设置高度为视口的高度
        }}
      >
        <Triangle
          visible={true}
          height="80"
          width="80"
          color="#4fa94d"
          ariaLabel="triangle-loading"
          // wrapperStyle和wrapperClass可根据需要进行调整
          wrapperStyle={{}}
          wrapperClass=""
        />
        <br />
        <p style={{ color: "green" }}>我是一个Ai男孩，站在DJ台，这里的气氛让我感觉很呀么很气派</p>
      </div>
    );
  };

  function insertComment(content, nickName, sayId) {
    const formData = new FormData();
    formData.append("content", content);
    formData.append("nickName", nickName);
    formData.append("sayId", sayId);

    fetch("http://localhost:1111/comment", {
      method: "POST",
      body: formData,
    })
      .then((response) => {
        const data = response.json();
        console.log(data);
      })
      .then((data) => console.log(data))
      .catch((error) => console.error("Error:", error));
  }

  function like(sayId) {
    const formData = new FormData();
    formData.append("sayId", sayId);

    fetch("http://localhost:1111/like", {
      method: "POST",
      body: formData,
    })
      .then((response) => {
        const data = response.json();
        console.log(data);
      })
      .then((data) => console.log(data))
      .catch((error) => console.error("Error:", error));
  }

  const getComments = async (say) => {
    try {
      const params = new URLSearchParams({
        sayId: say,
      });
      const response = await fetch(
        "http://localhost:1111/comments?" + params.toString()
      );
      const data = await response.json();
      console.log(data);
      setComments(data);
    } catch (error) {
      console.error("Failed to fetch data:", error);
    }
  };
  const [isLoading, setIsLoading] = useState(true);

  const fetchData2 = async () => {
    try {
      const response = await fetch("http://localhost:1111/says");
      const data = await response.json();

      // Format and dispatch the fetched data
      const formattedData = data.map((item) => ({
        content: item.content, // Assuming 'my' field expects the content of the message
        avatar: item.avatar,
        personName: item.personName,
        likeNum: item.likeNum, // Static value, assuming you might calculate or get it from elsewhere
        commentNum: item.commentNum, // Static value, assuming you might calculate or get it from elsewhere
        image: item.image,
        sayId: item.id,
      }));
      dispatch({ type: "barrages/saveDatas", payload: formattedData });
    } catch (error) {
      console.error("Failed to fetch data:", error);
    }
  };
  // Call the fetch function
  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch("http://localhost:1111/says");
        const data = await response.json();

        // Format and dispatch the fetched data
        const formattedData = data.map((item) => ({
          content: item.content, // Assuming 'my' field expects the content of the message
          avatar: item.avatar,
          personName: item.personName,
          likeNum: item.likeNum, // Static value, assuming you might calculate or get it from elsewhere
          commentNum: item.commentNum, // Static value, assuming you might calculate or get it from elsewhere
          image: item.image,
          sayId: item.id,
        }));
        setIsLoading(false);
        dispatch({
          type: "barrages/saveMessages",
          payload: shuffleArray(formattedData),
        });
        dispatch({ type: "barrages/saveDatas", payload: formattedData });
      } catch (error) {
        console.error("Failed to fetch data:", error);
      }
    };
    fetchData();
  }, [dispatch]);

  if (isLoading) {
    return Load();
  }

  return (
    <>
      <div className="container">
        <div className="barrage-container">
          {barrages.messages.map((message, index) => (
            <div className="barrage-message">
              <Card
                className="card"
                style={{ width: "30vh", marginTop: 16 }}
                actions={[
                  <div
                    className="dz"
                    onClick={() => {
                      like(message.sayId);
                      dispatch({
                        type: "barrages/addLike",
                        payload: { sayId: message.sayId },
                      });
                    }}
                  >
                    <div>
                      <HeartOutlined key="setting" />
                    </div>
                    <div className="dznum">{message.likeNum}</div>
                  </div>,
                  <div
                    className="dz"
                    onClick={() => {
                      showDrawer();
                      setSayId(message.sayId);
                      getComments(message.sayId);
                    }}
                  >
                    <div>
                      <MessageOutlined key="comment" />
                    </div>
                    <div className="dznum">{message.commentNum}</div>
                  </div>,
                ]}
                cover={<img alt="example" src={message.image} />}
              >
                <Skeleton loading={false} avatar active>
                  <Meta
                    avatar={<Avatar src={message.avatar} />}
                    title={message.personName}
                    description={message.content}
                  />
                </Skeleton>
              </Card>
            </div>
          ))}
        </div>

        <div className="list-bar">
          <br />
          <div
            id="scrollableDiv"
            style={{
              height: "99vh",
              overflow: "auto",
              padding: "0 12px",
              border: "0px solid blue",
              margin: "10",
            }}
          >
            <InfiniteScroll
              dataLength={100}
              // next={loadMoreData}
              // hasMore={data.length < 50}
              loader={<Skeleton avatar paragraph={{ rows: 1 }} active />}
              // endMessage={<Divider plain>It is all, nothing more 🤐</Divider>}
              scrollableTarget="scrollableDiv"
            >
              <center>
                <h3>秀句排行榜&nbsp;&nbsp;&nbsp;<Button shape="circle" onClick={() => {
                  fetchData2()
                }} icon={<RedoOutlined />}></Button></h3>
              </center>
              <List
                dataSource={barrages.data}
                renderItem={(item) => (
                  <List.Item key={item.content}>
                    <List.Item.Meta
                      avatar={<Avatar src={item.avatar} />}
                      title={item.personName}
                      description={item.content}
                    />
                    <div class="icon-with-text">
                      <div
                        class="icon-row"
                        onClick={() => {
                          like(item.sayId);
                          // fetchData2();
                          dispatch({
                            type: "barrages/addLikeData",
                            payload: { sayId: item.sayId },
                          });
                        }}
                      >
                        <HeartOutlined key="setting" />
                        <div class="text">{item.likeNum}</div>
                      </div>
                      <br />
                      <div
                        class="icon-row"
                        onClick={() => {
                          showDrawer();
                          setSayId(item.sayId);
                          getComments(item.sayId);
                        }}
                      >
                        <MessageOutlined className="dznum2" />
                        <div class="text">{item.commentNum}</div>
                      </div>
                    </div>
                  </List.Item>
                )}
              />
            </InfiniteScroll>
          </div>
        </div>
      </div>
      <Drawer
        title="评论列表"
        placement={placement}
        width={500}
        onClose={onClose}
        open={open}
        extra={
          <Space>
            <Button type="primary" onClick={showModal}>
              评论
            </Button>
          </Space>
        }
      >
        <List
          itemLayout="horizontal"
          dataSource={coms}
          renderItem={(item, index) => (
            <List.Item>
              <List.Item.Meta
                title={<a href="https://ant.design">{item.nickName + ":"}</a>}
                description={
                  <>
                    {item.content}
                    <br />
                    发布于：{item.recordTime}
                  </>
                }
              />
            </List.Item>
          )}
        />
      </Drawer>
      <Modal
        title="评论"
        open={openModal}
        onOk={() => {
          insertComment(com, nickNam, sayI);
          dispatch({ type: "barrages/addComment", payload: { sayId: sayI } });
          hideModal();
          getComments(sayI);
          setComment("");
        }}
        onCancel={hideModal}
        okText="确认"
        cancelText="取消"
      >
        <h4>
          你的昵称为：
          <Input
            style={{ width: "50%" }}
            defaultValue="孤独的狼"
            value={nickNam}
            onChange={(e) => setNickName(e.target.value)}
          />
          {/* <Button type="primary" style={{ margin: 10 }}>
            Ai随机
          </Button> */}
        </h4>
        <TextArea
          rows={4}
          value={com}
          onChange={(e) => setComment(e.target.value)}
        />
      </Modal>
    </>
  );
};

export default connect(({ barrages }) => ({
  barrages,
}))(Barrage);
