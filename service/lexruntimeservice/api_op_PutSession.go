// Code generated by smithy-go-codegen DO NOT EDIT.

package lexruntimeservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lexruntimeservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"io"
)

// Creates a new session or modifies an existing session with an Amazon Lex bot.
// Use this operation to enable your application to set the state of the bot. For
// more information, see Managing Sessions
// (https://docs.aws.amazon.com/lex/latest/dg/how-session-api.html).
func (c *Client) PutSession(ctx context.Context, params *PutSessionInput, optFns ...func(*Options)) (*PutSessionOutput, error) {
	stack := middleware.NewStack("PutSession", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpPutSessionMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	registerHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	addOpPutSessionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutSession(options.Region), middleware.Before)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     c.ServiceID(),
			OperationName: "PutSession",
			Err:           err,
		}
	}
	out := result.(*PutSessionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutSessionInput struct {
	// The message that Amazon Lex returns in the response can be either text or speech
	// based depending on the value of this field.
	//
	//     * If the value is text/plain;
	// charset=utf-8, Amazon Lex returns text in the response.
	//
	//     * If the value
	// begins with audio/, Amazon Lex returns speech in the response. Amazon Lex uses
	// Amazon Polly to generate the speech in the configuration that you specify. For
	// example, if you specify audio/mpeg as the value, Amazon Lex returns speech in
	// the MPEG format.
	//
	//     * If the value is audio/pcm, the speech is returned as
	// audio/pcm in 16-bit, little endian format.
	//
	//     * The following are the accepted
	// values:
	//
	//         * audio/mpeg
	//
	//         * audio/ogg
	//
	//         * audio/pcm
	//
	//
	// * audio/* (defaults to mpeg)
	//
	//         * text/plain; charset=utf-8
	Accept *string
	// The alias in use for the bot that contains the session data.
	BotAlias *string
	// The name of the bot that contains the session data.
	BotName *string
	// Sets the next action that the bot should take to fulfill the conversation.
	DialogAction *types.DialogAction
	// A summary of the recent intents for the bot. You can use the intent summary view
	// to set a checkpoint label on an intent and modify attributes of intents. You can
	// also use it to remove or add intent summary objects to the list. An intent that
	// you modify or add to the list must make sense for the bot. For example, the
	// intent name must be valid for the bot. You must provide valid values for:
	//
	//     *
	// intentName
	//
	//     * slot names
	//
	//     * slotToElict
	//
	// If you send the
	// recentIntentSummaryView parameter in a PutSession request, the contents of the
	// new summary view replaces the old summary view. For example, if a GetSession
	// request returns three intents in the summary view and you call PutSession with
	// one intent in the summary view, the next call to GetSession will only return one
	// intent.
	RecentIntentSummaryView []*types.IntentSummary
	// Map of key/value pairs representing the session-specific context information. It
	// contains application information passed between Amazon Lex and a client
	// application.
	SessionAttributes map[string]*string
	// The ID of the client application user. Amazon Lex uses this to identify a user's
	// conversation with your bot.
	UserId *string
}

type PutSessionOutput struct {
	// The audio version of the message to convey to the user.
	AudioStream io.ReadCloser
	// Content type as specified in the Accept HTTP header in the request.
	ContentType *string
	// * ConfirmIntent - Amazon Lex is expecting a "yes" or "no" response to confirm
	// the intent before fulfilling an intent.
	//
	//     * ElicitIntent - Amazon Lex wants
	// to elicit the user's intent.
	//
	//     * ElicitSlot - Amazon Lex is expecting the
	// value of a slot for the current intent.
	//
	//     * Failed - Conveys that the
	// conversation with the user has failed. This can happen for various reasons,
	// including the user does not provide an appropriate response to prompts from the
	// service, or if the Lambda function fails to fulfill the intent.
	//
	//     * Fulfilled
	// - Conveys that the Lambda function has sucessfully fulfilled the intent.
	//
	//     *
	// ReadyForFulfillment - Conveys that the client has to fulfill the intent.
	DialogState types.DialogState
	// The name of the current intent.
	IntentName *string
	// The next message that should be presented to the user.
	Message *string
	// The format of the response message. One of the following values:
	//
	//     *
	// PlainText - The message contains plain UTF-8 text.
	//
	//     * CustomPayload - The
	// message is a custom format for the client.
	//
	//     * SSML - The message contains
	// text formatted for voice output.
	//
	//     * Composite - The message contains an
	// escaped JSON object containing one or more messages from the groups that
	// messages were assigned to when the intent was created.
	MessageFormat types.MessageFormatType
	// Map of key/value pairs representing session-specific context information.
	// This value conforms to the media type: application/json
	SessionAttributes *string
	// A unique identifier for the session.
	SessionId *string
	// If the dialogState is ElicitSlot, returns the name of the slot for which Amazon
	// Lex is eliciting a value.
	SlotToElicit *string
	// Map of zero or more intent slots Amazon Lex detected from the user input during
	// the conversation. Amazon Lex creates a resolution list containing likely values
	// for a slot. The value that it returns is determined by the
	// valueSelectionStrategy selected when the slot type was created or updated. If
	// valueSelectionStrategy is set to ORIGINAL_VALUE, the value provided by the user
	// is returned, if the user value is similar to the slot values. If
	// valueSelectionStrategy is set to TOP_RESOLUTION Amazon Lex returns the first
	// value in the resolution list or, if there is no resolution list, null. If you
	// don't specify a valueSelectionStrategy the default is ORIGINAL_VALUE.
	// This value conforms to the media type: application/json
	Slots *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpPutSessionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpPutSession{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpPutSession{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutSession(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:         region,
		ServiceName:    "Lex Runtime Service",
		ServiceID:      "lexruntimeservice",
		EndpointPrefix: "lexruntimeservice",
		SigningName:    "lex",
		OperationName:  "PutSession",
	}
}
