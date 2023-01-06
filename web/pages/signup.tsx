import { Box, Divider, Text } from "@chakra-ui/react";
import {
  AppShell,
  Button,
  ButtonGroup,
  Card,
  CardBody,
  Field,
  FormLayout,
  FormStep,
  FormStepper,
  FormValue,
  Loader,
  NextButton,
  PrevButton,
  Property,
  PropertyList,
  StepForm,
  StepperCompleted,
  useModals,
  useSnackbar,
} from "@saas-ui/react";
import { Web3Address } from "@saas-ui/web3";
import Link from "next/link";
import { useState } from "react";
import * as Yup from "yup";

export default function SignUp() {
  const snackbar = useSnackbar();
  const modals = useModals();
  const [label, updateLabel] = useState("");
  const [account, setAccount] = useState<string | null>(null);
  const [credential, setCredential] = useState<PublicKeyCredential | null>(
    null
  );
  const schemas = {
    credential: Yup.object().shape({
      deviceLabel: Yup.string().required().label("Device Label"),
    }),
  };

  const yupResolver = (schema: any) => async (values: any) => {
    try {
      await schema.validate(values, { abortEarly: false });
      return {
        values,
        errors: {},
      };
    } catch (errors) {}
  };

  const onSubmit = (params: any) => {
    console.log(params);
    return new Promise((resolve) => {
      setTimeout(resolve, 1000);
    });
  };

  const createCredential = async (nextStep: () => void) => {
    const credential = await navigator.credentials.create({
      publicKey: {
        rp: {
          name: "ACME Corporation",
          id: "localhost",
        },
        user: {
          id: Uint8Array.from([1, 2, 3, 4, 5]),
          name: label,
          displayName: label,
        },
        challenge: Uint8Array.from([1, 2, 3, 4, 5]),
        pubKeyCredParams: [
          {
            type: "public-key",
            alg: -7,
          },
        ],
        timeout: 60000,
        attestation: "direct",
      },
    });
    setCredential(credential as PublicKeyCredential);
    snackbar.info("PassKey has been generated by your device.");
    nextStep();
  };

  const registerAccount = async (nextStep: () => void) => {
    modals.form({
      title: "Enter a Secure Recovery Passcode",
      body: (
        <>
          <Field
            isRequired
            key="pinCode"
            name="pinCode"
            label="Enter Passcode"
            type="pin"
            pinLength={6}
          />
          <Box key="pinConfirm" height="8" display="flex"></Box>
          <Field
            isRequired
            key="pinConfirm"
            name="pinConfirm"
            label="Re-enter Passcode"
            type="pin"
            pinLength={6}
          />
        </>
      ),
      onSubmit: async (values) => {
        return new Promise((resolve) => {
          if (values.pinCode !== values.pinConfirm) {
            snackbar.error(
              "Passcodes do not match, " +
                values.pinCode +
                " " +
                values.pinConfirm
            );
          } else {
            nextStep();
            modals.closeAll();
          }
          setTimeout(resolve, 1000);
        });
      },
    });
    const response = await fetch("/api/register", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        credential,
      }),
    });
    const { account } = await response.json();
    setAccount(account);

    nextStep();
  };

  return (
    <AppShell
      navbar={
        <>
          <Box
            as="header"
            alignItems="start"
            borderBottomWidth="1px"
            py="2"
            px="4"
          >
            <Link href="/">
              <Button variant="outline">Return Home</Button>
            </Link>
          </Box>
        </>
      }
    >
      <Box
        as="main"
        alignContent="center"
        marginLeft="20vw"
        marginRight="20vw"
        marginTop="15vh"
      >
        <Card
          isHoverable
          variant="solid"
          title="Sonr.ID"
          subtitle="Use the Vault MPC Protocol with Webauthn."
        >
          <CardBody>
            <StepForm
              defaultValues={{
                deviceLabel: "",
                credentialId: "",
                credentialType: "",
                credentialPublicKey: "",
                credentialResponse: "",
              }}
              onSubmit={onSubmit}
            >
              {({
                isFirstStep,
                isLastStep,
                isCompleted,
                nextStep,
                prevStep,
              }) => (
                <FormLayout>
                  <FormStepper orientation="vertical">
                    <FormStep
                      name="project"
                      title="Generate PassKey"
                      resolver={yupResolver(schemas.credential)}
                    >
                      <FormLayout>
                        <Field
                          isRequired
                          name="deviceLabel"
                          label="Device Label"
                          onInput={(event) => {
                            // Check if the input is a string
                            const str = event.target as HTMLInputElement;

                            // Get the value from the input
                            const value = str.value;
                            updateLabel(str.value);
                          }}
                        />
                        <Button
                          label="Open Webauthn"
                          onClick={() =>
                            label
                              ? createCredential(nextStep)
                              : snackbar.error("Please enter a device label")
                          }
                          variant="outline"
                        />
                      </FormLayout>
                    </FormStep>

                    <FormStep name="register" title="Register your Account">
                      <FormLayout>
                        <PropertyList>
                          <Property
                            label="Label"
                            value={label ? label : "No credential"}
                          />
                          <Property
                            label="Credential ID"
                            value={
                              <Web3Address
                                address={
                                  credential ? credential.id : "No credential"
                                }
                                startLength={credential ? 12 : 15}
                                endLength={credential ? 4 : 0}
                              />
                            }
                          />
                          <Property
                            label="Type"
                            value={
                              credential ? credential.type : "No credential"
                            }
                          />
                          <Property label="Source" value="WebAuthn" />
                        </PropertyList>
                        <Divider />
                        <ButtonGroup>
                          <Button
                            label="Continue"
                            onClick={() => registerAccount(nextStep)}
                          />
                          <Button label="Cancel" variant="ghost" />
                        </ButtonGroup>
                      </FormLayout>
                    </FormStep>

                    <FormStep name="faucet" title="Faucet Airdrop">
                      <FormLayout>
                        <Text>
                          Please confirm that your information is correct.
                        </Text>
                        <PropertyList>
                          <Property
                            label="Name"
                            value={<FormValue name="name" />}
                          />
                          <Property
                            label="Description"
                            value={<FormValue name="description" />}
                          />
                        </PropertyList>
                        <ButtonGroup>
                          <NextButton />
                          <PrevButton variant="ghost" />
                        </ButtonGroup>
                      </FormLayout>
                    </FormStep>
                    <FormStep name="confirm" title="Broadcast Transaction">
                      <FormLayout>
                        <Text>
                          Please confirm that your information is correct.
                        </Text>
                        <PropertyList>
                          <Property
                            label="Name"
                            value={<FormValue name="name" />}
                          />
                          <Property
                            label="Description"
                            value={<FormValue name="description" />}
                          />
                        </PropertyList>
                        <ButtonGroup>
                          <NextButton />
                          <PrevButton variant="ghost" />
                        </ButtonGroup>
                      </FormLayout>
                    </FormStep>

                    <StepperCompleted>
                      <Loader>
                        We are setting up your project, just a moment...
                      </Loader>
                    </StepperCompleted>
                  </FormStepper>
                </FormLayout>
              )}
            </StepForm>
          </CardBody>
        </Card>
      </Box>
    </AppShell>
  );
}
