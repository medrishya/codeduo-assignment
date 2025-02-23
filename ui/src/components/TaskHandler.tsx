import React, { useState, useEffect } from "react";
import {
  List,
  Button,
  Row,
  Col,
  Modal,
  Divider,
  Pagination,
  message,
} from "antd";
import { createTask, deleteTask, fetchTasks, updateTask } from "../api/api";
import { Task } from "../types";
import TaskForm from "./TaskForm";
import TaskItem from "./TaskItem";

const TaskManager: React.FC = () => {
  const [tasks, setTasks] = useState<Task[]>([]);
  const [currentTask, setCurrentTask] = useState<Task | null>(null);
  const [isModalVisible, setIsModalVisible] = useState(false);
  const [pagination, setPagination] = useState({
    currentPage: 1,
    pageSize: 5,
    total: 0,
  });
  const [loadingList, setLoadingList] = useState(false);
  const [loadingSave, setLoadingSave] = useState(false);
  const [messageApi, contextHolder] = message.useMessage();

  useEffect(() => {
    const loadTasks = async () => {
      setLoadingList(true);
      try {
        const data = await fetchTasks(
          pagination.currentPage,
          pagination.pageSize
        );
        setTasks(data.tasks || []);
        setPagination((prev) => ({ ...prev, total: data.total }));
      } catch (error) {
        messageApi.error("Failed to load tasks.");
      } finally {
        setLoadingList(false);
      }
    };

    loadTasks();
  }, [pagination.currentPage, pagination.pageSize]);

  const showModal = (task?: Task) => {
    setCurrentTask(task || null);
    setIsModalVisible(true);
  };

  const handleSubmit = async (values: { name: string; status: string }) => {
    setLoadingSave(true);
    try {
      const updatedTask = currentTask
        ? await updateTask(currentTask.id, values)
        : await createTask(values);
      setTasks((prev) => {
        const newTasks = currentTask
          ? prev.map((task) =>
              task.id === currentTask.id ? updatedTask : task
            )
          : [...prev, updatedTask];
        messageApi.success(
          `Task ${currentTask ? "updated" : "added"} successfully!`
        );
        return newTasks;
      });
      setIsModalVisible(false);
    } catch {
      messageApi.error("Failed to save task!");
    } finally {
      setLoadingSave(false);
    }
  };

  const handleDelete = async (id: number) => {
    try {
      await deleteTask(id);
      setTasks((prev) => prev.filter((task) => task.id !== id));
      messageApi.success("Task deleted successfully!");
    } catch {
      messageApi.error("Failed to delete task!");
    }
  };

  const toggleStatus = async (task: Task) => {
    try {
      const updatedTask = await updateTask(task.id, {
        status: task.status === "completed" ? "pending" : "completed",
      });
      setTasks((prev) => prev.map((t) => (t.id === task.id ? updatedTask : t)));
    } catch {
      messageApi.error("Failed to update task status!");
    }
  };

  const handlePageChange = (currentPage: number, pageSize?: number) => {
    setPagination((prev) => ({
      ...prev,
      currentPage,
      pageSize: pageSize || prev.pageSize,
    }));
  };

  return (
    <div style={{ padding: 20 }}>
      {contextHolder}
      <Row justify="end" style={{ marginBottom: 16 }}>
        <Col>
          <Button type="primary" onClick={() => showModal()}>
            Add Task
          </Button>
        </Col>
      </Row>
      <List
        bordered
        loading={loadingList}
        dataSource={tasks}
        renderItem={(task, index) => (
          <TaskItem
            task={task}
            index={index}
            onEdit={() => showModal(task)}
            onToggle={() => toggleStatus(task)}
            onDelete={handleDelete}
          />
        )}
        style={{ maxWidth: 600, margin: "0 auto" }}
      />
      <Divider />
      <Pagination
        current={pagination.currentPage}
        pageSize={pagination.pageSize}
        total={pagination.total}
        onChange={handlePageChange}
        showSizeChanger
        pageSizeOptions={["5", "10", "15"]}
      />
      <Modal
        title={currentTask ? "Edit Task" : "Add Task"}
        open={isModalVisible}
        footer={null}
        onCancel={() => setIsModalVisible(false)}
        width={300}
      >
        <TaskForm
          initialValues={currentTask}
          onSubmit={handleSubmit}
          loading={loadingSave}
        />
      </Modal>
    </div>
  );
};

export default TaskManager;
