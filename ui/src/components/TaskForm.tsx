import React from "react";
import { Form, Input, Select, Button } from "antd";
import { TaskFormProps } from "../types";

const options = [
  { value: "pending", label: "Pending" },
  { value: "completed", label: "Completed" },
];

const TaskForm: React.FC<TaskFormProps> = ({
  initialValues,
  onSubmit,
  loading,
}) => {
  const [form] = Form.useForm();

  React.useEffect(() => {
    if (initialValues) {
      form.setFieldsValue({
        name: initialValues.name,
        status: initialValues.status,
      });
    } else {
      form.resetFields();
    }
  }, [initialValues, form]);

  return (
    <Form form={form} layout="vertical" onFinish={onSubmit}>
      <Form.Item
        name="name"
        label="Task Name"
        rules={[{ required: true, message: "Please enter the task name." }]}
      >
        <Input />
      </Form.Item>
      <Form.Item
        name="status"
        label="Task Status"
        rules={[{ required: true, message: "Please select task status." }]}
      >
        <Select placeholder="Select status">
          {options.map((option) => (
            <Select.Option key={option.value} value={option.value}>
              {option.label}
            </Select.Option>
          ))}
        </Select>
      </Form.Item>
      <Form.Item>
        <Button type="primary" htmlType="submit" loading={loading}>
          Save Task
        </Button>
      </Form.Item>
    </Form>
  );
};

export default TaskForm;
