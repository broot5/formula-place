interface PageHeaderProps {
  pageName: string;
  action?: React.ReactNode;
}

export function PageHeader({ pageName, action }: PageHeaderProps) {
  return (
    <div className="flex items-center justify-between">
      <h1 className="text-2xl font-bold tracking-tight">{pageName}</h1>
      {action}
    </div>
  );
}
