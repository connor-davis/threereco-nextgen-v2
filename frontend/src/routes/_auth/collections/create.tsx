import { postApiTransactionsMutation } from '@/api-client/@tanstack/react-query.gen';
import { useMutation } from '@tanstack/react-query';
import {
  ErrorComponent,
  Link,
  createFileRoute,
  useRouter,
} from '@tanstack/react-router';
import {
  ArrowLeftIcon,
  BrickWallIcon,
  CheckIcon,
  InfoIcon,
  LoaderCircleIcon,
  TrashIcon,
  UserIcon,
} from 'lucide-react';
import { useForm } from 'react-hook-form';

import { zodResolver } from '@hookform/resolvers/zod';
import { toast } from 'sonner';
import z from 'zod';

import {
  type CreateTransaction,
  type ErrorResponse,
  type Material,
  type User,
  getApiMaterials,
  getApiUsers,
} from '@/api-client';
import { zCreateTransaction } from '@/api-client/zod.gen';
import PermissionGuard from '@/components/guards/permission';
import { Alert, AlertDescription, AlertTitle } from '@/components/ui/alert';
import { Badge } from '@/components/ui/badge';
import { Button } from '@/components/ui/button';
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/components/ui/card';
import { Checkbox } from '@/components/ui/checkbox';
import { DebounceInput } from '@/components/ui/debounce-input';
import { DebounceNumberInput } from '@/components/ui/debounce-number-input';
import { Form } from '@/components/ui/form';
import { Label } from '@/components/ui/label';
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from '@/components/ui/popover';
import {
  Stepper,
  StepperContent,
  StepperIndicator,
  StepperItem,
  StepperNav,
  StepperPanel,
  StepperSeparator,
  StepperTitle,
  StepperTrigger,
} from '@/components/ui/stepper';
import { apiClient } from '@/lib/utils';

export const Route = createFileRoute('/_auth/collections/create')({
  component: () => (
    <PermissionGuard value="transactions.create" isPage={true}>
      <RouteComponent />
    </PermissionGuard>
  ),
  validateSearch: z.object({
    materialsPage: z.coerce.number().default(1),
    materialsSearch: z.string().default(''),
    accountsPage: z.coerce.number().default(1),
    accountsSearch: z.string().default(''),
    step: z.coerce.number().default(1),
  }),
  pendingComponent: () => (
    <div className="flex flex-col w-full h-full items-center justify-center">
      <Label className="text-muted-foreground">Loading collections...</Label>
    </div>
  ),
  errorComponent: ({ error }: { error: Error | ErrorResponse }) => {
    if ('error' in error) {
      // Render a custom error message
      return (
        <div className="flex flex-col w-full h-full items-center justify-center">
          <Alert variant="destructive" className="w-full max-w-lg">
            <AlertTitle>{error.error}</AlertTitle>
            <AlertDescription>{error.message}</AlertDescription>
          </Alert>
        </div>
      );
    }

    // Fallback to the default ErrorComponent
    return <ErrorComponent error={error} />;
  },
  wrapInSuspense: true,
  loaderDeps: ({
    search: {
      materialsPage,
      materialsSearch,
      accountsPage,
      accountsSearch,
      step,
    },
  }) => ({
    materialsPage,
    materialsSearch,
    accountsPage,
    accountsSearch,
    step,
  }),
  loader: async ({
    deps: { materialsPage, materialsSearch, accountsPage, accountsSearch },
  }) => {
    const { data: products } = await getApiMaterials({
      client: apiClient,
      query: {
        page: materialsPage,
        search: materialsSearch,
        limit: 10,
      },
      throwOnError: true,
    });

    const { data: accounts } = await getApiUsers({
      client: apiClient,
      query: {
        page: accountsPage,
        search: accountsSearch,
        limit: 10,
        type: 'business',
      },
      throwOnError: true,
    });

    return {
      materials: (products.items ?? []) as Array<Material>,
      materialsPageDetails: products.pageDetails ?? {},
      accounts: (accounts.items ?? []) as Array<User>,
      accountsPageDetails: accounts.pageDetails ?? {},
    };
  },
});

function RouteComponent() {
  const router = useRouter();
  const { materialsPage, materialsSearch, accountsPage, accountsSearch, step } =
    Route.useLoaderDeps();
  const { materials, materialsPageDetails, accounts, accountsPageDetails } =
    Route.useLoaderData();

  const createForm = useForm<CreateTransaction>({
    resolver: zodResolver(zCreateTransaction),
    defaultValues: {
      sellerId: undefined,
      buyerId: undefined,
    },
  });

  const createTransaction = useMutation({
    ...postApiTransactionsMutation({
      client: apiClient,
    }),
    onError: (error: ErrorResponse) =>
      toast.error(error.error, {
        description: error.message,
        duration: 2000,
      }),
    onSuccess: () => {
      toast.success('Success', {
        description: 'The collection has been created successfully.',
        duration: 2000,
      });

      router.navigate({
        to: '/collections/create',
        search: {
          materialsPage,
          materialsSearch,
          accountsPage,
          accountsSearch,
          step: 4,
        },
      });

      createForm.reset();

      return router.invalidate();
    },
  });

  return (
    <div className="flex flex-col w-full h-full bg-popover border-t p-3 gap-3 overflow-hidden">
      <div className="flex items-center justify-between w-full h-auto">
        <div className="flex items-center gap-3">
          <Link to="/collections">
            <Button variant="ghost" size="icon">
              <ArrowLeftIcon className="size-4" />
            </Button>
          </Link>

          <Label className="text-lg">Create Collection</Label>
        </div>
        <div className="flex items-center gap-3"></div>
      </div>

      <Form {...createForm}>
        <form
          onSubmit={createForm.handleSubmit(
            (values) =>
              createTransaction.mutate({
                body: values,
              }),
            () =>
              router.navigate({
                to: '/collections/create',
                search: {
                  materialsPage,
                  materialsSearch,
                  accountsPage,
                  accountsSearch,
                  step: 1,
                },
              })
          )}
          className="flex flex-col w-full h-full overflow-hidden"
        >
          <Stepper
            value={step}
            onValueChange={(step) =>
              router.navigate({
                to: '/collections/create',
                search: {
                  materialsPage,
                  materialsSearch,
                  accountsPage,
                  accountsSearch,
                  step,
                },
              })
            }
            indicators={{
              completed: <CheckIcon className="size-4" />,
              loading: <LoaderCircleIcon className="size-4 animate-spin" />,
            }}
            className="flex flex-col w-full h-full gap-5 overflow-hidden"
          >
            <StepperNav className="flex gap-3 h-32">
              <StepperItem step={1} className="relative flex-1 items-start">
                <StepperTrigger
                  className="flex flex-col items-start justify-center gap-2.5 grow"
                  asChild
                >
                  <StepperIndicator className="size-8 border-2 data-[state=completed]:text-white data-[state=completed]:bg-primary data-[state=inactive]:bg-transparent data-[state=inactive]:border-border data-[state=inactive]:text-muted-foreground">
                    <BrickWallIcon className="size-4" />
                  </StepperIndicator>
                  <div className="flex flex-col items-start gap-1">
                    <div className="text-[10px] font-semibold uppercase text-muted-foreground">
                      Step 1
                    </div>
                    <StepperTitle className="text-start text-base font-semibold group-data-[state=inactive]/step:text-muted-foreground">
                      Collection Materials
                    </StepperTitle>
                    <div>
                      <Badge
                        variant="secondary"
                        className="hidden group-data-[state=active]/step:inline-flex"
                      >
                        In Progress
                      </Badge>
                      <Badge
                        variant="default"
                        className="hidden group-data-[state=completed]/step:inline-flex"
                      >
                        Completed
                      </Badge>
                      <Badge
                        variant="outline"
                        className="hidden group-data-[state=inactive]/step:inline-flex text-muted-foreground"
                      >
                        Pending
                      </Badge>
                    </div>
                  </div>
                </StepperTrigger>
                <StepperSeparator className="absolute top-4 inset-x-0 start-9 m-0 group-data-[orientation=horizontal]/stepper-nav:w-[calc(100%-2rem)] group-data-[orientation=horizontal]/stepper-nav:flex-none  group-data-[state=completed]/step:bg-primary" />
              </StepperItem>

              <StepperItem
                step={2}
                className="relative flex-1 items-start"
                loading={createTransaction.isPending}
              >
                <StepperTrigger
                  className="flex flex-col items-start justify-center gap-2.5 grow"
                  asChild
                >
                  <StepperIndicator className="size-8 border-2 data-[state=completed]:text-white data-[state=completed]:bg-primary data-[state=inactive]:bg-transparent data-[state=inactive]:border-border data-[state=inactive]:text-muted-foreground">
                    <UserIcon className="size-4" />
                  </StepperIndicator>
                  <div className="flex flex-col items-start gap-1">
                    <div className="text-[10px] font-semibold uppercase text-muted-foreground">
                      Step 2
                    </div>
                    <StepperTitle className="text-start text-base font-semibold group-data-[state=inactive]/step:text-muted-foreground">
                      Collection Collector
                    </StepperTitle>
                    <div>
                      <Badge
                        variant="secondary"
                        className="hidden group-data-[state=active]/step:inline-flex"
                      >
                        In Progress
                      </Badge>
                      <Badge
                        variant="default"
                        className="hidden group-data-[state=completed]/step:inline-flex"
                      >
                        Completed
                      </Badge>
                      <Badge
                        variant="outline"
                        className="hidden group-data-[state=inactive]/step:inline-flex text-muted-foreground"
                      >
                        Pending
                      </Badge>
                    </div>
                  </div>
                </StepperTrigger>
                <StepperSeparator className="absolute top-4 inset-x-0 start-9 m-0 group-data-[orientation=horizontal]/stepper-nav:w-[calc(100%-2rem)] group-data-[orientation=horizontal]/stepper-nav:flex-none  group-data-[state=completed]/step:bg-primary" />
              </StepperItem>

              <StepperItem
                step={3}
                className="relative flex-1 items-start"
                loading={createTransaction.isPending}
              >
                <StepperTrigger
                  className="flex flex-col items-start justify-center gap-2.5 grow"
                  asChild
                >
                  <StepperIndicator className="size-8 border-2 data-[state=completed]:text-white data-[state=completed]:bg-primary data-[state=inactive]:bg-transparent data-[state=inactive]:border-border data-[state=inactive]:text-muted-foreground">
                    <InfoIcon className="size-4" />
                  </StepperIndicator>
                  <div className="flex flex-col items-start gap-1">
                    <div className="text-[10px] font-semibold uppercase text-muted-foreground">
                      Step 3
                    </div>
                    <StepperTitle className="text-start text-base font-semibold group-data-[state=inactive]/step:text-muted-foreground">
                      Collection Overview
                    </StepperTitle>
                    <div>
                      <Badge
                        variant="secondary"
                        className="hidden group-data-[state=active]/step:inline-flex"
                      >
                        In Progress
                      </Badge>
                      <Badge
                        variant="default"
                        className="hidden group-data-[state=completed]/step:inline-flex"
                      >
                        Completed
                      </Badge>
                      <Badge
                        variant="outline"
                        className="hidden group-data-[state=inactive]/step:inline-flex text-muted-foreground"
                      >
                        Pending
                      </Badge>
                    </div>
                  </div>
                </StepperTrigger>
                <StepperSeparator className="absolute top-4 inset-x-0 start-9 m-0 group-data-[orientation=horizontal]/stepper-nav:w-[calc(100%-2rem)] group-data-[orientation=horizontal]/stepper-nav:flex-none  group-data-[state=completed]/step:bg-primary" />
              </StepperItem>

              <StepperItem
                step={4}
                className="relative items-start"
                loading={createTransaction.isPending}
              >
                <StepperTrigger
                  className="flex flex-col items-start justify-center gap-2.5"
                  asChild
                >
                  <StepperIndicator className="size-8 border-2 data-[state=completed]:text-white data-[state=completed]:bg-primary data-[state=inactive]:bg-transparent data-[state=inactive]:border-border data-[state=inactive]:text-muted-foreground">
                    <CheckIcon className="size-4" />
                  </StepperIndicator>
                  <div className="flex flex-col items-start gap-1">
                    <div className="text-[10px] font-semibold uppercase text-muted-foreground"></div>
                    <StepperTitle className="text-start text-base font-semibold group-data-[state=inactive]/step:text-muted-foreground">
                      Collection Created
                    </StepperTitle>
                    <div>
                      <Badge
                        variant="secondary"
                        className="hidden group-data-[state=active]/step:inline-flex"
                      >
                        In Progress
                      </Badge>
                      <Badge
                        variant="default"
                        className="hidden group-data-[state=completed]/step:inline-flex"
                      >
                        Completed
                      </Badge>
                      <Badge
                        variant="outline"
                        className="hidden group-data-[state=inactive]/step:inline-flex text-muted-foreground"
                      >
                        Pending
                      </Badge>
                    </div>
                  </div>
                </StepperTrigger>
              </StepperItem>
            </StepperNav>

            <StepperPanel className="flex flex-col w-full h-full overflow-hidden">
              <StepperContent
                value={1}
                className="flex flex-col w-full h-full overflow-hidden gap-3"
              >
                <div className="flex flex-col w-full h-full overflow-hidden gap-3">
                  <div className="flex items-center justify-between w-full h-auto gap-3">
                    <div className="flex items-center gap-3">
                      <Label>Added Materials</Label>
                    </div>

                    <div className="flex items-center gap-3">
                      <Popover>
                        <PopoverTrigger>
                          <Button>Available Materials</Button>
                        </PopoverTrigger>
                        <PopoverContent className="w-full lg:w-120">
                          <div className="flex flex-col w-full h-full overflow-hidden gap-3">
                            <div className="flex flex-col w-full h-auto gap-3">
                              <DebounceInput
                                type="text"
                                placeholder="Search materials..."
                                className="w-full"
                                defaultValue={materialsSearch}
                                onChange={(e) => {
                                  const search = e.target.value;

                                  router.navigate({
                                    to: '/collections/create',
                                    search: {
                                      materialsPage,
                                      materialsSearch: search,
                                      accountsPage,
                                      accountsSearch,
                                      step,
                                    },
                                  });
                                }}
                              />
                            </div>

                            <div className="flex flex-col w-full h-full overflow-y-auto gap-3">
                              {materials.map((material) => (
                                <Label className="hover:bg-accent flex items-center justify-between gap-3 rounded-lg border p-3">
                                  <Checkbox
                                    id="toggle-2"
                                    checked={(
                                      createForm.watch().materials ?? []
                                    )
                                      .map((_material) => _material.materialId)
                                      .includes(material.id)}
                                    onCheckedChange={(checked) => {
                                      if (checked) {
                                        createForm.setValue('materials', [
                                          ...(createForm.getValues()
                                            .materials ?? []),
                                          {
                                            materialId: material.id,
                                            weight: 0,
                                            value: 0,
                                          },
                                        ]);
                                      } else {
                                        createForm.setValue('materials', [
                                          ...(
                                            createForm.getValues().materials ??
                                            []
                                          ).filter(
                                            (_material) =>
                                              _material.materialId !==
                                              material.id
                                          ),
                                        ]);
                                      }
                                    }}
                                    className="data-[state=checked]:border-primary data-[state=checked]:bg-primary data-[state=checked]:text-white"
                                  />

                                  <div className="flex items-center justify-between w-full h-auto gap-3">
                                    <div className="flex flex-col">
                                      <Label>{material.name}</Label>
                                    </div>

                                    <div className="flex items-center gap-2">
                                      <Badge>GW {material.gwCode}</Badge>
                                      <Badge>
                                        tC02e {material.carbonFactor}
                                      </Badge>
                                    </div>
                                  </div>
                                </Label>
                              ))}
                            </div>

                            <div className="flex items-center justify-end w-full">
                              <Label className="text-xs text-muted-foreground">
                                Page {materialsPage} of{' '}
                                {materialsPageDetails.pages}
                              </Label>

                              <Link
                                to="/collections/create"
                                search={{
                                  materialsPage:
                                    materialsPageDetails.previousPage,
                                  materialsSearch,
                                  accountsPage,
                                  accountsSearch,
                                  step,
                                }}
                                disabled={
                                  materialsPage ===
                                  materialsPageDetails.previousPage
                                }
                                reloadDocument={false}
                              >
                                <Button
                                  variant="outline"
                                  className="ml-3"
                                  disabled={
                                    materialsPage ===
                                    materialsPageDetails.previousPage
                                  }
                                >
                                  Previous
                                </Button>
                              </Link>
                              <Link
                                to="/collections/create"
                                search={{
                                  materialsPage: materialsPageDetails.nextPage,
                                  materialsSearch,
                                  accountsPage,
                                  accountsSearch,
                                  step,
                                }}
                                disabled={
                                  materialsPage ===
                                  materialsPageDetails.nextPage
                                }
                                reloadDocument={false}
                              >
                                <Button
                                  variant="outline"
                                  className="ml-1"
                                  disabled={
                                    materialsPage ===
                                    materialsPageDetails.nextPage
                                  }
                                >
                                  Next
                                </Button>
                              </Link>
                            </div>
                          </div>
                        </PopoverContent>
                      </Popover>
                    </div>
                  </div>

                  <div className="flex flex-col w-full h-full overflow-y-auto gap-3">
                    {(createForm.getValues().materials ?? []).length === 0 ? (
                      <div className="flex flex-col w-full h-full items-center justify-center gap-3">
                        <Label className="text-muted-foreground">
                          No materials added
                        </Label>
                      </div>
                    ) : (
                      createForm
                        .getValues()
                        .materials?.map((material) => {
                          const foundMaterial = materials.find(
                            (_material) => _material.id === material.materialId
                          );

                          return { ...material, foundMaterial };
                        })
                        .filter((material) => material.foundMaterial)
                        .map((material) => (
                          <div className="flex flex-col w-full h-auto gap-5 bg-accent rounded-lg border p-3">
                            <div className="flex items-center justify-between w-full h-auto gap-3">
                              <div className="flex flex-col">
                                <Label>{material.foundMaterial!.name}</Label>
                              </div>

                              <div className="flex items-center gap-2">
                                <Badge>
                                  GW {material.foundMaterial!.gwCode}
                                </Badge>
                                <Badge>
                                  tC02e {material.foundMaterial!.carbonFactor}
                                </Badge>

                                <Button
                                  variant="outline"
                                  size="icon"
                                  onClick={() =>
                                    createForm.setValue('materials', [
                                      ...(
                                        createForm.getValues().materials ?? []
                                      ).filter(
                                        (_material) =>
                                          _material.materialId !==
                                          material.materialId
                                      ),
                                    ])
                                  }
                                >
                                  <TrashIcon className="size-4" />
                                </Button>
                              </div>
                            </div>

                            <div className="flex flex-col w-full h-auto gap-3">
                              <Label>Weight</Label>
                              <DebounceNumberInput
                                decimalScale={2}
                                fixedDecimalScale
                                className="w-full"
                                value={material.weight}
                                onValueChange={(value) =>
                                  createForm.setValue(
                                    'materials',
                                    (
                                      createForm.getValues().materials ?? []
                                    ).map((_material) => {
                                      if (
                                        _material.materialId ===
                                        material.materialId
                                      ) {
                                        return {
                                          ..._material,
                                          weight: value ?? 0,
                                        };
                                      }
                                      return _material;
                                    })
                                  )
                                }
                              />
                              <Label className="text-sm text-muted-foreground">
                                The materials weight (kg)
                              </Label>
                            </div>

                            <div className="flex flex-col w-full h-auto gap-3">
                              <Label>Value</Label>
                              <DebounceNumberInput
                                decimalScale={2}
                                fixedDecimalScale
                                className="w-full"
                                value={material.value}
                                onValueChange={(value) =>
                                  createForm.setValue(
                                    'materials',
                                    (
                                      createForm.getValues().materials ?? []
                                    ).map((_material) => {
                                      if (
                                        _material.materialId ===
                                        material.materialId
                                      ) {
                                        return {
                                          ..._material,
                                          value: value ?? 0,
                                        };
                                      }
                                      return _material;
                                    })
                                  )
                                }
                              />
                              <Label className="text-sm text-muted-foreground">
                                The materials value (R)
                              </Label>
                            </div>
                          </div>
                        ))
                    )}
                  </div>
                </div>

                <div className="grid grid-cols-2 w-full h-auto gap-5 items-center">
                  <Link to="/collections" className="w-full">
                    <Button type="button" variant="outline" className="w-full">
                      Cancel
                    </Button>
                  </Link>
                  <Button
                    type="button"
                    className="w-full"
                    onClick={() =>
                      router.navigate({
                        to: '/collections/create',
                        search: {
                          materialsPage,
                          materialsSearch,
                          accountsPage,
                          accountsSearch,
                          step: 2,
                        },
                      })
                    }
                  >
                    Continue
                  </Button>
                </div>
              </StepperContent>

              <StepperContent
                value={2}
                className="flex flex-col w-full h-full gap-3"
              >
                <div className="flex flex-col w-full h-full overflow-hidden gap-3">
                  <div className="flex items-center justify-between w-full h-auto gap-3">
                    <div className="flex items-center gap-3">
                      <Label>Available Collectors</Label>
                    </div>

                    <div className="flex items-center gap-3">
                      <DebounceInput
                        type="text"
                        placeholder="Search collectors..."
                        className="w-64"
                        defaultValue={materialsSearch}
                        onChange={(e) => {
                          const search = e.target.value;

                          router.navigate({
                            to: '/collections/create',
                            search: {
                              materialsPage,
                              materialsSearch,
                              accountsPage,
                              accountsSearch: search,
                              step,
                            },
                          });
                        }}
                      />
                    </div>
                  </div>

                  <div className="flex flex-col w-full h-full overflow-y-auto gap-3">
                    {accounts.map((account) => (
                      <Label className="hover:bg-accent flex items-center justify-between gap-3 rounded-lg border p-3">
                        <Checkbox
                          id="toggle-2"
                          checked={
                            (createForm.watch().sellerId ?? []) === account.id
                          }
                          onCheckedChange={(checked) => {
                            if (checked) {
                              createForm.setValue('sellerId', account.id);
                            } else {
                              createForm.setValue('sellerId', '');
                            }
                          }}
                          className="data-[state=checked]:border-primary data-[state=checked]:bg-primary data-[state=checked]:text-white"
                        />

                        <div className="flex items-center justify-between w-full h-auto gap-3">
                          {account.name && (
                            <div className="flex flex-col">
                              <Label className="text-sm">{account.name}</Label>
                              <Label className="text-sm text-muted-foreground">
                                {account.email}
                              </Label>
                            </div>
                          )}

                          {account.email && (
                            <div className="flex flex-col">
                              <Label className="text-sm">{account.email}</Label>
                            </div>
                          )}

                          {!account.email && account.phone && (
                            <div className="flex flex-col">
                              <Label className="text-sm">{account.phone}</Label>
                            </div>
                          )}
                        </div>
                      </Label>
                    ))}
                  </div>

                  <div className="flex items-center justify-end w-full">
                    <Label className="text-xs text-muted-foreground">
                      Page {accountsPage} of {accountsPageDetails.pages}
                    </Label>

                    <Link
                      to="/collections/create"
                      search={{
                        materialsPage,
                        materialsSearch,
                        accountsPage: accountsPageDetails.previousPage,
                        accountsSearch,
                        step,
                      }}
                      disabled={
                        accountsPage === accountsPageDetails.previousPage
                      }
                      reloadDocument={false}
                    >
                      <Button
                        variant="outline"
                        className="ml-3"
                        disabled={
                          accountsPage === accountsPageDetails.previousPage
                        }
                      >
                        Previous
                      </Button>
                    </Link>
                    <Link
                      to="/collections/create"
                      search={{
                        materialsPage,
                        materialsSearch,
                        accountsPage: accountsPageDetails.nextPage,
                        accountsSearch,
                        step,
                      }}
                      disabled={accountsPage === accountsPageDetails.nextPage}
                      reloadDocument={false}
                    >
                      <Button
                        variant="outline"
                        className="ml-1"
                        disabled={accountsPage === accountsPageDetails.nextPage}
                      >
                        Next
                      </Button>
                    </Link>
                  </div>
                </div>

                <div className="grid grid-cols-2 w-full h-auto gap-5 items-center">
                  <Button
                    type="button"
                    variant="outline"
                    className="w-full"
                    onClick={() =>
                      router.navigate({
                        to: '/collections/create',
                        search: {
                          materialsPage,
                          materialsSearch,
                          accountsPage,
                          accountsSearch,
                          step: 1,
                        },
                      })
                    }
                  >
                    Back
                  </Button>
                  <Button
                    type="button"
                    className="w-full"
                    onClick={() =>
                      router.navigate({
                        to: '/collections/create',
                        search: {
                          materialsPage,
                          materialsSearch,
                          accountsPage,
                          accountsSearch,
                          step: 3,
                        },
                      })
                    }
                  >
                    Continue
                  </Button>
                </div>
              </StepperContent>

              <StepperContent
                value={3}
                className="flex flex-col w-full h-full gap-3"
              >
                <div className="grid grid-cols-2 w-full h-full overflow-hidden gap-3">
                  <div className="flex flex-col w-full h-auto gap-3">
                    <Card className="border-dashed">
                      <CardHeader>
                        <CardTitle>Collector</CardTitle>
                        <CardDescription>
                          Review the collector details.
                        </CardDescription>
                      </CardHeader>
                      <CardContent className="flex flex-col w-full h-full overflow-hidden">
                        <div className="flex flex-col w-full h-auto gap-5">
                          <div className="flex flex-col w-full h-auto gap-3">
                            <Label className="text-muted-foreground">
                              Name
                            </Label>
                            <Label>
                              {accounts.find(
                                (account) =>
                                  account.id === createForm.getValues().sellerId
                              )?.name ?? 'No name'}
                            </Label>
                          </div>

                          <div className="flex flex-col w-full h-auto gap-3">
                            <Label className="text-muted-foreground">
                              Contact Info
                            </Label>
                            <Label>
                              {accounts.find(
                                (account) =>
                                  account.id === createForm.getValues().sellerId
                              )?.email
                                ? accounts.find(
                                    (account) =>
                                      account.id ===
                                      createForm.getValues().sellerId
                                  )?.email
                                : accounts.find(
                                      (account) =>
                                        account.id ===
                                        createForm.getValues().sellerId
                                    )?.phone
                                  ? accounts.find(
                                      (account) =>
                                        account.id ===
                                        createForm.getValues().sellerId
                                    )?.phone
                                  : 'No contact info'}
                            </Label>
                          </div>
                        </div>
                      </CardContent>
                    </Card>

                    <Card className="border-dashed">
                      <CardHeader>
                        <CardTitle>Added Materials</CardTitle>
                        <CardDescription>
                          Review the materials added to the collection.
                        </CardDescription>
                      </CardHeader>
                      <CardContent className="flex flex-col w-full h-full overflow-hidden">
                        <div className="flex flex-col w-full h-autogap-3">
                          {(createForm.getValues().materials ?? []).length ===
                          0 ? (
                            <Label className="text-muted-foreground">
                              No materials added
                            </Label>
                          ) : (
                            createForm
                              .getValues()
                              .materials?.map((material) => {
                                const foundMaterial = materials.find(
                                  (_material) =>
                                    _material.id === material.materialId
                                );

                                if (!foundMaterial) return null;

                                return (
                                  <Label
                                    key={foundMaterial.id}
                                    className="hover:bg-accent flex items-center justify-between gap-3 rounded-lg border p-3"
                                  >
                                    <div className="flex items-center justify-between w-full h-auto gap-3">
                                      <div className="flex flex-col w-full h-auto gap-10">
                                        <div className="flex items-center justify-between w-full h-auto gap-3">
                                          <div className="flex flex-col">
                                            <Label>{foundMaterial.name}</Label>
                                          </div>

                                          <div className="flex items-center gap-2">
                                            <Badge>
                                              GW {foundMaterial.gwCode}
                                            </Badge>
                                            <Badge>
                                              tC02e {foundMaterial.carbonFactor}
                                            </Badge>
                                          </div>
                                        </div>

                                        <div className="flex flex-col w-full h-auto gap-5">
                                          <div className="flex items-center w-full h-auto justify-between gap-3">
                                            <Label>Value</Label>
                                            <Label className="text-xs text-muted-foreground">
                                              {`${new Intl.NumberFormat(
                                                'en-ZA',
                                                {
                                                  style: 'currency',
                                                  currency: 'ZAR',
                                                }
                                              ).format(material.value ?? 0)}`}
                                            </Label>
                                          </div>

                                          <div className="flex items-center w-full h-auto justify-between gap-3">
                                            <Label>Weight</Label>
                                            <Label className="text-xs text-muted-foreground">
                                              {`${new Intl.NumberFormat(
                                                'en-ZA',
                                                {
                                                  style: 'unit',
                                                  unit: 'kilogram',
                                                }
                                              ).format(material.weight ?? 0)}`}
                                            </Label>
                                          </div>
                                        </div>
                                      </div>
                                    </div>
                                  </Label>
                                );
                              })
                          )}
                        </div>
                      </CardContent>
                    </Card>
                  </div>

                  <Card className="border-dashed">
                    <CardHeader>
                      <CardTitle>Collection Details</CardTitle>
                      <CardDescription>
                        Review the collection details before creating.
                      </CardDescription>
                    </CardHeader>
                    <CardContent className="flex flex-col w-full h-full gap-3 justify-start items-start">
                      <div className="flex flex-col items-start justify-start w-full h-auto gap-3">
                        <Label className="text-muted-foreground">Amount</Label>
                        <Label>
                          {new Intl.NumberFormat('en-ZA', {
                            style: 'currency',
                            currency: 'ZAR',
                          }).format(
                            (createForm
                              .getValues()
                              .materials?.map((material) => material.value ?? 0)
                              .reduce((a, b) => a + b, 0) ?? 0) *
                              (createForm
                                .getValues()
                                .materials?.map(
                                  (material) => material.weight ?? 0
                                )
                                .reduce((a, b) => a + b, 0) ?? 0)
                          )}
                        </Label>
                        <Label className="text-sm text-muted-foreground">
                          This will be the collections amount.
                        </Label>
                      </div>

                      <div className="flex flex-col items-start justify-start w-full h-auto gap-3">
                        <Label className="text-muted-foreground">Weight</Label>
                        <Label>
                          {new Intl.NumberFormat('en-ZA', {
                            style: 'unit',
                            unit: 'kilogram',
                          }).format(
                            createForm
                              .getValues()
                              .materials?.map(
                                (material) => material.weight ?? 0
                              )
                              .reduce((a, b) => a + b, 0) ?? 0
                          )}
                        </Label>
                        <Label className="text-sm text-muted-foreground">
                          This will be the collections weight.
                        </Label>
                      </div>
                    </CardContent>
                    <CardFooter>
                      <div className="grid grid-cols-2 w-full h-auto gap-5 items-center">
                        <Button
                          type="button"
                          variant="outline"
                          className="w-full"
                          onClick={() =>
                            router.navigate({
                              to: '/collections/create',
                              search: {
                                materialsPage,
                                materialsSearch,
                                accountsPage,
                                accountsSearch,
                                step: 2,
                              },
                            })
                          }
                        >
                          Back
                        </Button>
                        <Button className="w-full">Create Collection</Button>
                      </div>
                    </CardFooter>
                  </Card>
                </div>
              </StepperContent>

              <StepperContent
                value={4}
                className="flex flex-col w-full h-full gap-3"
              >
                <div className="flex flex-col w-full h-full gap-5">
                  <div className="flex flex-col w-full h-auto items-center justify-center gap-3">
                    <p className="text-lg font-semibold">
                      Collection created successfully!
                    </p>
                    <p className="text-sm text-muted-foreground">
                      You can view your collection in the collection list.
                    </p>
                  </div>

                  <Button
                    className="w-full"
                    onClick={() =>
                      router.navigate({
                        to: '/collections/create',
                        search: {
                          materialsPage,
                          materialsSearch,
                          accountsPage,
                          accountsSearch,
                          step: 1,
                        },
                      })
                    }
                  >
                    Create New Collection
                  </Button>
                </div>
              </StepperContent>
            </StepperPanel>
          </Stepper>
        </form>
      </Form>
    </div>
  );
}
