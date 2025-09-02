import { deleteApiRolesByIdMutation } from '@/api-client/@tanstack/react-query.gen';
import { useMutation } from '@tanstack/react-query';
import { useRouter } from '@tanstack/react-router';
import { TrashIcon } from 'lucide-react';
import { type ReactNode, useState } from 'react';

import { toast } from 'sonner';

import type { ErrorResponse } from '@/api-client';
import { apiClient } from '@/lib/utils';

import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from '../ui/alert-dialog';
import { Button } from '../ui/button';
import { Input } from '../ui/input';

export default function DeleteRoleByIdDialog({
  id,
  name,
  children,
}: {
  id: string;
  name: string;
  children?: ReactNode;
}) {
  const router = useRouter();

  const [confirmationValue, setConfirmationValue] = useState<string>('');

  const deleteRole = useMutation({
    ...deleteApiRolesByIdMutation({
      client: apiClient,
    }),
    onError: (error: ErrorResponse) =>
      toast.error(error.error, {
        description: error.message,
        duration: 2000,
      }),
    onSuccess: () => {
      toast.success('Success', {
        description: 'The role has been deleted successfully.',
        duration: 2000,
      });

      return router.invalidate();
    },
  });

  return (
    <AlertDialog>
      <AlertDialogTrigger>
        {children ?? (
          <Button variant="outline" size="icon">
            <TrashIcon />
          </Button>
        )}
      </AlertDialogTrigger>
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>Are you absolutely sure?</AlertDialogTitle>
          <AlertDialogDescription>
            This action will permanently remove the role from the organization.
            Please type <strong>{name}</strong> to confirm.
          </AlertDialogDescription>
        </AlertDialogHeader>

        <Input
          type="text"
          autoComplete="off"
          value={confirmationValue}
          onChange={(e) => setConfirmationValue(e.target.value)}
          placeholder={name}
        />

        <AlertDialogFooter>
          <AlertDialogCancel>Cancel</AlertDialogCancel>
          <AlertDialogAction
            disabled={!confirmationValue || confirmationValue !== name}
            onClick={() =>
              deleteRole.mutate({
                path: {
                  id,
                },
              })
            }
          >
            Delete Role
          </AlertDialogAction>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  );
}
