export interface Task {
  id: number;
  name: string;
  status: string;
}

export interface TaskItemProps {
  task: Task;
  index: number;
  onEdit: () => void;
  onToggle: () => void;
  onDelete: (id: number) => void;
}

export interface TaskFormProps {
  initialValues: Task | null;
  onSubmit: (values: { name: string; status: string }) => void;
  loading: boolean;
}
