import React from "react";
import { List, Button, Switch, Popconfirm } from "antd";
import { DeleteFilled, EditOutlined } from "@ant-design/icons";
import { TaskItemProps } from "../types";

const TaskItem: React.FC<TaskItemProps> = ({
  task,
  index,
  onEdit,
  onDelete,
  onToggle,
}) => (
  <List.Item
    actions={[
      <Switch
        checked={task.status === "completed"}
        onChange={() => onToggle()}
        checkedChildren="completed"
        unCheckedChildren={task.status}
      />,
      <Button
        type="link"
        shape="circle"
        icon={<EditOutlined />}
        onClick={onEdit}
      />,
      <Popconfirm
        title="Are you sure to delete this task?"
        onConfirm={() => onDelete(task.id)}
        okText="Yes"
        cancelText="No"
      >
        <Button danger type="link" icon={<DeleteFilled />} />
      </Popconfirm>,
    ]}
  >
    <div>
      <strong>{index + 1}.</strong> {task.name}
    </div>
  </List.Item>
);

export default TaskItem;
