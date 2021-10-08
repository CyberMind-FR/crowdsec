// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/crowdsecurity/crowdsec/pkg/database/ent/alert"
	"github.com/crowdsecurity/crowdsec/pkg/database/ent/bouncer"
	"github.com/crowdsecurity/crowdsec/pkg/database/ent/decision"
	"github.com/crowdsecurity/crowdsec/pkg/database/ent/event"
	"github.com/crowdsecurity/crowdsec/pkg/database/ent/machine"
	"github.com/crowdsecurity/crowdsec/pkg/database/ent/meta"
	"github.com/crowdsecurity/crowdsec/pkg/database/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	alertFields := schema.Alert{}.Fields()
	_ = alertFields
	// alertDescCreatedAt is the schema descriptor for created_at field.
	alertDescCreatedAt := alertFields[0].Descriptor()
	// alert.DefaultCreatedAt holds the default value on creation for the created_at field.
	alert.DefaultCreatedAt = alertDescCreatedAt.Default.(func() time.Time)
	// alertDescUpdatedAt is the schema descriptor for updated_at field.
	alertDescUpdatedAt := alertFields[1].Descriptor()
	// alert.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	alert.DefaultUpdatedAt = alertDescUpdatedAt.Default.(func() time.Time)
	// alertDescBucketId is the schema descriptor for bucketId field.
	alertDescBucketId := alertFields[3].Descriptor()
	// alert.DefaultBucketId holds the default value on creation for the bucketId field.
	alert.DefaultBucketId = alertDescBucketId.Default.(string)
	// alertDescMessage is the schema descriptor for message field.
	alertDescMessage := alertFields[4].Descriptor()
	// alert.DefaultMessage holds the default value on creation for the message field.
	alert.DefaultMessage = alertDescMessage.Default.(string)
	// alertDescEventsCount is the schema descriptor for eventsCount field.
	alertDescEventsCount := alertFields[5].Descriptor()
	// alert.DefaultEventsCount holds the default value on creation for the eventsCount field.
	alert.DefaultEventsCount = alertDescEventsCount.Default.(int32)
	// alertDescStartedAt is the schema descriptor for startedAt field.
	alertDescStartedAt := alertFields[6].Descriptor()
	// alert.DefaultStartedAt holds the default value on creation for the startedAt field.
	alert.DefaultStartedAt = alertDescStartedAt.Default.(func() time.Time)
	// alertDescStoppedAt is the schema descriptor for stoppedAt field.
	alertDescStoppedAt := alertFields[7].Descriptor()
	// alert.DefaultStoppedAt holds the default value on creation for the stoppedAt field.
	alert.DefaultStoppedAt = alertDescStoppedAt.Default.(func() time.Time)
	// alertDescSimulated is the schema descriptor for simulated field.
	alertDescSimulated := alertFields[21].Descriptor()
	// alert.DefaultSimulated holds the default value on creation for the simulated field.
	alert.DefaultSimulated = alertDescSimulated.Default.(bool)
	bouncerFields := schema.Bouncer{}.Fields()
	_ = bouncerFields
	// bouncerDescCreatedAt is the schema descriptor for created_at field.
	bouncerDescCreatedAt := bouncerFields[0].Descriptor()
	// bouncer.DefaultCreatedAt holds the default value on creation for the created_at field.
	bouncer.DefaultCreatedAt = bouncerDescCreatedAt.Default.(func() time.Time)
	// bouncerDescUpdatedAt is the schema descriptor for updated_at field.
	bouncerDescUpdatedAt := bouncerFields[1].Descriptor()
	// bouncer.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	bouncer.DefaultUpdatedAt = bouncerDescUpdatedAt.Default.(func() time.Time)
	// bouncerDescIPAddress is the schema descriptor for ip_address field.
	bouncerDescIPAddress := bouncerFields[5].Descriptor()
	// bouncer.DefaultIPAddress holds the default value on creation for the ip_address field.
	bouncer.DefaultIPAddress = bouncerDescIPAddress.Default.(string)
	// bouncerDescUntil is the schema descriptor for until field.
	bouncerDescUntil := bouncerFields[8].Descriptor()
	// bouncer.DefaultUntil holds the default value on creation for the until field.
	bouncer.DefaultUntil = bouncerDescUntil.Default.(func() time.Time)
	// bouncerDescLastPull is the schema descriptor for last_pull field.
	bouncerDescLastPull := bouncerFields[9].Descriptor()
	// bouncer.DefaultLastPull holds the default value on creation for the last_pull field.
	bouncer.DefaultLastPull = bouncerDescLastPull.Default.(func() time.Time)
	decisionFields := schema.Decision{}.Fields()
	_ = decisionFields
	// decisionDescCreatedAt is the schema descriptor for created_at field.
	decisionDescCreatedAt := decisionFields[0].Descriptor()
	// decision.DefaultCreatedAt holds the default value on creation for the created_at field.
	decision.DefaultCreatedAt = decisionDescCreatedAt.Default.(func() time.Time)
	// decisionDescUpdatedAt is the schema descriptor for updated_at field.
	decisionDescUpdatedAt := decisionFields[1].Descriptor()
	// decision.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	decision.DefaultUpdatedAt = decisionDescUpdatedAt.Default.(func() time.Time)
	// decisionDescSimulated is the schema descriptor for simulated field.
	decisionDescSimulated := decisionFields[13].Descriptor()
	// decision.DefaultSimulated holds the default value on creation for the simulated field.
	decision.DefaultSimulated = decisionDescSimulated.Default.(bool)
	eventFields := schema.Event{}.Fields()
	_ = eventFields
	// eventDescCreatedAt is the schema descriptor for created_at field.
	eventDescCreatedAt := eventFields[0].Descriptor()
	// event.DefaultCreatedAt holds the default value on creation for the created_at field.
	event.DefaultCreatedAt = eventDescCreatedAt.Default.(func() time.Time)
	// eventDescUpdatedAt is the schema descriptor for updated_at field.
	eventDescUpdatedAt := eventFields[1].Descriptor()
	// event.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	event.DefaultUpdatedAt = eventDescUpdatedAt.Default.(func() time.Time)
	// eventDescSerialized is the schema descriptor for serialized field.
	eventDescSerialized := eventFields[3].Descriptor()
	// event.SerializedValidator is a validator for the "serialized" field. It is called by the builders before save.
	event.SerializedValidator = eventDescSerialized.Validators[0].(func(string) error)
	machineFields := schema.Machine{}.Fields()
	_ = machineFields
	// machineDescCreatedAt is the schema descriptor for created_at field.
	machineDescCreatedAt := machineFields[0].Descriptor()
	// machine.DefaultCreatedAt holds the default value on creation for the created_at field.
	machine.DefaultCreatedAt = machineDescCreatedAt.Default.(func() time.Time)
	// machineDescUpdatedAt is the schema descriptor for updated_at field.
	machineDescUpdatedAt := machineFields[1].Descriptor()
	// machine.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	machine.DefaultUpdatedAt = machineDescUpdatedAt.Default.(func() time.Time)
	// machineDescLastPush is the schema descriptor for last_push field.
	machineDescLastPush := machineFields[2].Descriptor()
	// machine.DefaultLastPush holds the default value on creation for the last_push field.
	machine.DefaultLastPush = machineDescLastPush.Default.(func() time.Time)
	// machineDescScenarios is the schema descriptor for scenarios field.
	machineDescScenarios := machineFields[6].Descriptor()
	// machine.ScenariosValidator is a validator for the "scenarios" field. It is called by the builders before save.
	machine.ScenariosValidator = machineDescScenarios.Validators[0].(func(string) error)
	// machineDescIsValidated is the schema descriptor for isValidated field.
	machineDescIsValidated := machineFields[8].Descriptor()
	// machine.DefaultIsValidated holds the default value on creation for the isValidated field.
	machine.DefaultIsValidated = machineDescIsValidated.Default.(bool)
	metaFields := schema.Meta{}.Fields()
	_ = metaFields
	// metaDescCreatedAt is the schema descriptor for created_at field.
	metaDescCreatedAt := metaFields[0].Descriptor()
	// meta.DefaultCreatedAt holds the default value on creation for the created_at field.
	meta.DefaultCreatedAt = metaDescCreatedAt.Default.(func() time.Time)
	// metaDescUpdatedAt is the schema descriptor for updated_at field.
	metaDescUpdatedAt := metaFields[1].Descriptor()
	// meta.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	meta.DefaultUpdatedAt = metaDescUpdatedAt.Default.(func() time.Time)
	// metaDescValue is the schema descriptor for value field.
	metaDescValue := metaFields[3].Descriptor()
	// meta.ValueValidator is a validator for the "value" field. It is called by the builders before save.
	meta.ValueValidator = metaDescValue.Validators[0].(func(string) error)
}
