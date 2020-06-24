package ns


/*
#cgo CFLAGS: -x objective-c -fno-objc-arc
#cgo LDFLAGS: -framework Foundation
#pragma clang diagnostic ignored "-Wformat-security"

#import <Foundation/Foundation.h>

void*
NSObject_MutableCopyWithZone(void* zone) {
	NSObject* ret;
	@autoreleasepool {
		ret = [NSObject mutableCopyWithZone:zone];
	}
	return ret;

}
void*
NSObject_New() {
	NSObject* ret;
	@autoreleasepool {
		ret = [NSObject new];
	}
	return ret;

}
BOOL
NSObject_IsSubclassOfClass(void* aClass) {
	BOOL ret;
	@autoreleasepool {
		ret = [NSObject isSubclassOfClass:aClass];
	}
	return ret;

}
void*
selectorFromString(char *s) {
	return NSSelectorFromString([NSString stringWithUTF8String:s]);
}

void*
NSObject_InstanceMethodSignatureForSelector(void* aSelector) {
	NSMethodSignature* ret;
	@autoreleasepool {
		ret = [NSObject instanceMethodSignatureForSelector:aSelector];
	}
	return ret;

}
void*
NSObject_Alloc() {
	return [NSObject alloc];
}
BOOL
NSObject_AccessInstanceVariablesDirectly() {
	BOOL ret;
	@autoreleasepool {
		ret = [NSObject accessInstanceVariablesDirectly];
	}
	return ret;

}
void
NSObject_SetVersion(NSInteger aVersion) {
	@autoreleasepool {
		[NSObject setVersion:aVersion];
	}
}
BOOL
NSObject_UseStoredAccessor() {
	BOOL ret;
	@autoreleasepool {
		ret = [NSObject useStoredAccessor];
	}
	return ret;

}
void
NSObject_Load() {
	@autoreleasepool {
		[NSObject load];
	}
}
void* _Nonnull
NSObject_KeyPathsForValuesAffectingValueForKey(void* key) {
	NSSet* _Nonnull ret;
	@autoreleasepool {
		ret = [NSObject keyPathsForValuesAffectingValueForKey:key];
	}
	return ret;

}
NSInteger
NSObject_Version() {
	NSInteger ret;
	@autoreleasepool {
		ret = [NSObject version];
	}
	return ret;

}
BOOL
NSObject_ConformsToProtocol(void* protocol) {
	BOOL ret;
	@autoreleasepool {
		ret = [NSObject conformsToProtocol:protocol];
	}
	return ret;

}
void
NSObject_CancelPreviousPerformRequestsWithTarget(void* aTarget) {
	@autoreleasepool {
		[NSObject cancelPreviousPerformRequestsWithTarget:aTarget];
	}
}
void
NSObject_CancelPreviousPerformRequestsWithTargetSelector(void* aTarget, void* aSelector, void* anArgument) {
	@autoreleasepool {
		[NSObject cancelPreviousPerformRequestsWithTarget:aTarget selector:aSelector object:anArgument];
	}
}
BOOL
NSObject_InstancesRespondToSelector(void* aSelector) {
	BOOL ret;
	@autoreleasepool {
		ret = [NSObject instancesRespondToSelector:aSelector];
	}
	return ret;

}
void*
NSObject_DebugDescription() {
	NSString* ret;
	@autoreleasepool {
		ret = [NSObject debugDescription];
	}
	return ret;

}
void*
NSObject_Class() {
	Class ret;
	@autoreleasepool {
		ret = [NSObject class];
	}
	return ret;

}
void* _Nonnull
NSObject_ClassFallbacksForKeyedArchiver() {
	NSArray* _Nonnull ret;
	@autoreleasepool {
		ret = [NSObject classFallbacksForKeyedArchiver];
	}
	return ret;

}
void
NSObject_SetKeys(void* keys, void* dependentKey) {
	@autoreleasepool {
		[NSObject setKeys:keys triggerChangeNotificationsForDependentKey:dependentKey];
	}
}
BOOL
NSObject_ResolveClassMethod(void* sel) {
	BOOL ret;
	@autoreleasepool {
		ret = [NSObject resolveClassMethod:sel];
	}
	return ret;

}
void*
NSObject_Superclass() {
	Class ret;
	@autoreleasepool {
		ret = [NSObject superclass];
	}
	return ret;

}
BOOL
NSObject_AutomaticallyNotifiesObserversForKey(void* key) {
	BOOL ret;
	@autoreleasepool {
		ret = [NSObject automaticallyNotifiesObserversForKey:key];
	}
	return ret;

}
void*
NSObject_AllocWithZone(void* zone) {
	return [NSObject allocWithZone:zone];
}
BOOL
NSObject_ResolveInstanceMethod(void* sel) {
	BOOL ret;
	@autoreleasepool {
		ret = [NSObject resolveInstanceMethod:sel];
	}
	return ret;

}
void* _Nonnull
NSObject_ClassForKeyedUnarchiver() {
	Class _Nonnull ret;
	@autoreleasepool {
		ret = [NSObject classForKeyedUnarchiver];
	}
	return ret;

}
void*
NSObject_Description() {
	NSString* ret;
	@autoreleasepool {
		ret = [NSObject description];
	}
	return ret;

}
void*
NSObject_CopyWithZone(void* zone) {
	NSObject* ret;
	@autoreleasepool {
		ret = [NSObject copyWithZone:zone];
	}
	return ret;

}
NSUInteger
NSObject_Hash() {
	NSUInteger ret;
	@autoreleasepool {
		ret = [NSObject hash];
	}
	return ret;

}
void*
NSObject_inst_Description(void* o) {
	NSString* ret;
	@autoreleasepool {
		ret = [(NSObject*)o description];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void*
NSObject_inst_MethodSignatureForSelector(void* o, void* aSelector) {
	NSMethodSignature* ret;
	@autoreleasepool {
		ret = [(NSObject*)o methodSignatureForSelector:aSelector];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void
NSObject_inst_SetValuesForKeysWithDictionary(void* o, void* keyedValues) {
	@autoreleasepool {
		[(NSObject*)o setValuesForKeysWithDictionary:keyedValues];
	}
}
void
NSObject_inst_ObserveValueForKeyPath(void* o, void* keyPath, void* object, void* change, void* context) {
	@autoreleasepool {
		[(NSObject*)o observeValueForKeyPath:keyPath ofObject:object change:change context:context];
	}
}
void* _Nullable
NSObject_inst_StoredValueForKey(void* o, void* key) {
	NSObject* _Nullable ret;
	@autoreleasepool {
		ret = [(NSObject*)o storedValueForKey:key];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void* _Nullable
NSObject_inst_ClassForArchiver(void* o) {
	Class _Nullable ret;
	@autoreleasepool {
		ret = [(NSObject*)o classForArchiver];
	}
	return ret;

}
void* _Nullable
NSObject_inst_ReplacementObjectForPortCoder(void* o, void* coder) {
	NSObject* _Nullable ret;
	@autoreleasepool {
		ret = [(NSObject*)o replacementObjectForPortCoder:coder];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void
NSObject_inst_URLResourceDataDidBecomeAvailable(void* o, void* sender, void* newBytes) {
	@autoreleasepool {
		[(NSObject*)o URL:sender resourceDataDidBecomeAvailable:newBytes];
	}
}
void
NSObject_inst_URLResourceDidFailLoadingWithReason(void* o, void* sender, void* reason) {
	@autoreleasepool {
		[(NSObject*)o URL:sender resourceDidFailLoadingWithReason:reason];
	}
}
void* _Nullable
NSObject_inst_ValueForKeyPath(void* o, void* keyPath) {
	NSObject* _Nullable ret;
	@autoreleasepool {
		ret = [(NSObject*)o valueForKeyPath:keyPath];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void
NSObject_inst_UnableToSetNilForKey(void* o, void* key) {
	@autoreleasepool {
		[(NSObject*)o unableToSetNilForKey:key];
	}
}
BOOL
NSObject_inst_ScriptingIsGreaterThanOrEqualTo(void* o, void* object) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSObject*)o scriptingIsGreaterThanOrEqualTo:object];
	}
	return ret;

}
void* _Nonnull
NSObject_inst_ClassForPortCoder(void* o) {
	Class _Nonnull ret;
	@autoreleasepool {
		ret = [(NSObject*)o classForPortCoder];
	}
	return ret;

}
void
NSObject_inst_InsertValueInPropertyWithKey(void* o, void* value, void* key) {
	@autoreleasepool {
		[(NSObject*)o insertValue:value inPropertyWithKey:key];
	}
}
void
NSObject_inst_InsertValueAtIndex(void* o, void* value, NSUInteger index, void* key) {
	@autoreleasepool {
		[(NSObject*)o insertValue:value atIndex:index inPropertyWithKey:key];
	}
}
void
NSObject_inst_WillChange(void* o, NSKeyValueChange changeKind, void* indexes, void* key) {
	@autoreleasepool {
		[(NSObject*)o willChange:changeKind valuesAtIndexes:indexes forKey:key];
	}
}
BOOL
NSObject_inst_AttemptRecoveryFromErrorOptionIndex(void* o, void* error, NSUInteger recoveryOptionIndex) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSObject*)o attemptRecoveryFromError:error optionIndex:recoveryOptionIndex];
	}
	return ret;

}
void
NSObject_inst_AttemptRecoveryFromErrorOptionIndexDelegate(void* o, void* error, NSUInteger recoveryOptionIndex, void* delegate, void* didRecoverSelector, void* contextInfo) {
	@autoreleasepool {
		[(NSObject*)o attemptRecoveryFromError:error optionIndex:recoveryOptionIndex delegate:delegate didRecoverSelector:didRecoverSelector contextInfo:contextInfo];
	}
}
BOOL
NSObject_inst_ScriptingIsLessThan(void* o, void* object) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSObject*)o scriptingIsLessThan:object];
	}
	return ret;

}
void
NSObject_inst_Release(void* o) {
	@autoreleasepool {
		[(NSObject*)o release];
	}
}
void* _Nullable
NSObject_inst_NewScriptingObjectOfClass(void* o, void* objectClass, void* key, void* contentsValue, void* properties) {
	NSObject* _Nullable ret;
	@autoreleasepool {
		ret = [(NSObject*)o newScriptingObjectOfClass:objectClass forValueForKey:key withContentsValue:contentsValue properties:properties];
	}
	return ret;

}
void* _Nonnull
NSObject_inst_MutableSetValueForKeyPath(void* o, void* keyPath) {
	NSMutableSet* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSObject*)o mutableSetValueForKeyPath:keyPath];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void
NSObject_inst_WillChangeValueForKey(void* o, void* key) {
	@autoreleasepool {
		[(NSObject*)o willChangeValueForKey:key];
	}
}
void
NSObject_inst_WillChangeValueForKeyWithSetMutation(void* o, void* key, NSKeyValueSetMutationKind mutationKind, void* objects) {
	@autoreleasepool {
		[(NSObject*)o willChangeValueForKey:key withSetMutation:mutationKind usingObjects:objects];
	}
}
void
NSObject_inst_Dealloc(void* o) {
	@autoreleasepool {
		[(NSObject*)o dealloc];
	}
}
void*
NSObject_inst_Retain(void* o) {
	NSObject* ret;
	@autoreleasepool {
		ret = [(NSObject*)o retain];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void
NSObject_inst_DidChange(void* o, NSKeyValueChange changeKind, void* indexes, void* key) {
	@autoreleasepool {
		[(NSObject*)o didChange:changeKind valuesAtIndexes:indexes forKey:key];
	}
}
BOOL
NSObject_inst_IsLike(void* o, void* object) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSObject*)o isLike:object];
	}
	return ret;

}
void* _Nonnull
NSObject_inst_MutableSetValueForKey(void* o, void* key) {
	NSMutableSet* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSObject*)o mutableSetValueForKey:key];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
BOOL
NSObject_inst_IsKindOfClass(void* o, void* aClass) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSObject*)o isKindOfClass:aClass];
	}
	return ret;

}
void
NSObject_inst_ReplaceValueAtIndex(void* o, NSUInteger index, void* key, void* value) {
	@autoreleasepool {
		[(NSObject*)o replaceValueAtIndex:index inPropertyWithKey:key withValue:value];
	}
}
BOOL
NSObject_inst_IsEqualTo(void* o, void* object) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSObject*)o isEqualTo:object];
	}
	return ret;

}
void* _Nonnull
NSObject_inst_DictionaryWithValuesForKeys(void* o, void* keys) {
	NSDictionary* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSObject*)o dictionaryWithValuesForKeys:keys];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
NSUInteger
NSObject_inst_RetainCount(void* o) {
	NSUInteger ret;
	@autoreleasepool {
		ret = [(NSObject*)o retainCount];
	}
	return ret;

}
void
NSObject_inst_ForwardInvocation(void* o, void* anInvocation) {
	@autoreleasepool {
		[(NSObject*)o forwardInvocation:anInvocation];
	}
}
void
NSObject_inst_TakeValueForKey(void* o, void* value, void* key) {
	@autoreleasepool {
		[(NSObject*)o takeValue:value forKey:key];
	}
}
void
NSObject_inst_TakeValueForKeyPath(void* o, void* value, void* keyPath) {
	@autoreleasepool {
		[(NSObject*)o takeValue:value forKeyPath:keyPath];
	}
}
void*
NSObject_inst_Copy(void* o) {
	NSObject* ret;
	@autoreleasepool {
		ret = [(NSObject*)o copy];
	}
	return ret;

}
void
NSObject_inst_URLResourceDidCancelLoading(void* o, void* sender) {
	@autoreleasepool {
		[(NSObject*)o URLResourceDidCancelLoading:sender];
	}
}
BOOL
NSObject_inst_FileManagerShouldProceedAfterError(void* o, void* fm, void* errorInfo) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSObject*)o fileManager:fm shouldProceedAfterError:errorInfo];
	}
	return ret;

}
void
NSObject_inst_FileManagerWillProcessPath(void* o, void* fm, void* path) {
	@autoreleasepool {
		[(NSObject*)o fileManager:fm willProcessPath:path];
	}
}
BOOL
NSObject_inst_ConformsToProtocol(void* o, void* aProtocol) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSObject*)o conformsToProtocol:aProtocol];
	}
	return ret;

}
void* _Nullable
NSObject_inst_ClassForKeyedArchiver(void* o) {
	Class _Nullable ret;
	@autoreleasepool {
		ret = [(NSObject*)o classForKeyedArchiver];
	}
	return ret;

}
void
NSObject_inst_SetValueForKey(void* o, void* value, void* key) {
	@autoreleasepool {
		[(NSObject*)o setValue:value forKey:key];
	}
}
void
NSObject_inst_SetValueForKeyPath(void* o, void* value, void* keyPath) {
	@autoreleasepool {
		[(NSObject*)o setValue:value forKeyPath:keyPath];
	}
}
void
NSObject_inst_SetValueForUndefinedKey(void* o, void* value, void* key) {
	@autoreleasepool {
		[(NSObject*)o setValue:value forUndefinedKey:key];
	}
}
BOOL
NSObject_inst_IsMemberOfClass(void* o, void* aClass) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSObject*)o isMemberOfClass:aClass];
	}
	return ret;

}
void
NSObject_inst_SetObservationInfo(void* o, void* observationInfo) {
	@autoreleasepool {
		[(NSObject*)o setObservationInfo:observationInfo];
	}
}
void* _Nonnull
NSObject_inst_ToManyRelationshipKeys(void* o) {
	NSArray* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSObject*)o toManyRelationshipKeys];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void
NSObject_inst_RemoveObserverForKeyPath(void* o, void* observer, void* keyPath) {
	@autoreleasepool {
		[(NSObject*)o removeObserver:observer forKeyPath:keyPath];
	}
}
void
NSObject_inst_RemoveObserverForKeyPathContext(void* o, void* observer, void* keyPath, void* context) {
	@autoreleasepool {
		[(NSObject*)o removeObserver:observer forKeyPath:keyPath context:context];
	}
}
void* _Nonnull
NSObject_inst_MutableOrderedSetValueForKeyPath(void* o, void* keyPath) {
	NSMutableOrderedSet* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSObject*)o mutableOrderedSetValueForKeyPath:keyPath];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void
NSObject_inst_AddObserver(void* o, void* observer, void* keyPath, NSKeyValueObservingOptions options, void* context) {
	@autoreleasepool {
		[(NSObject*)o addObserver:observer forKeyPath:keyPath options:options context:context];
	}
}
void
NSObject_inst_DoesNotRecognizeSelector(void* o, void* aSelector) {
	@autoreleasepool {
		[(NSObject*)o doesNotRecognizeSelector:aSelector];
	}
}
BOOL
NSObject_inst_ValidateValueForKey(void* o, void** ioValue, void* inKey, void** outError) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSObject*)o validateValue:(id _Nullable* _Nonnull)ioValue forKey:inKey error:(NSError* _Nullable* _Nullable)outError];
		for(int i=0;i<1;i++) {
			if(ioValue[i] == 0) { break; }
			[(id)ioValue[i] retain];
		}
	
	
		for(int i=0;i<1;i++) {
			if(outError[i] == 0) { break; }
			[(id)outError[i] retain];
		}
	
	}
	return ret;

}
BOOL
NSObject_inst_ValidateValueForKeyPath(void* o, void** ioValue, void* inKeyPath, void** outError) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSObject*)o validateValue:(id _Nullable* _Nonnull)ioValue forKeyPath:inKeyPath error:(NSError* _Nullable* _Nullable)outError];
		for(int i=0;i<1;i++) {
			if(ioValue[i] == 0) { break; }
			[(id)ioValue[i] retain];
		}
	
	
		for(int i=0;i<1;i++) {
			if(outError[i] == 0) { break; }
			[(id)outError[i] retain];
		}
	
	}
	return ret;

}
void*
NSObject_inst_Self(void* o) {
	NSObject* ret;
	@autoreleasepool {
		ret = [(NSObject*)o self];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void* _Nullable
NSObject_inst_ScriptingProperties(void* o) {
	NSDictionary* _Nullable ret;
	@autoreleasepool {
		ret = [(NSObject*)o scriptingProperties];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void*
NSObject_inst_MutableCopy(void* o) {
	NSObject* ret;
	@autoreleasepool {
		ret = [(NSObject*)o mutableCopy];
	}
	return ret;

}
void* _Nullable
NSObject_inst_ValueAtIndex(void* o, NSUInteger index, void* key) {
	NSObject* _Nullable ret;
	@autoreleasepool {
		ret = [(NSObject*)o valueAtIndex:index inPropertyWithKey:key];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void* _Nonnull
NSObject_inst_ClassDescription(void* o) {
	NSClassDescription* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSObject*)o classDescription];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void
NSObject_inst_TakeStoredValue(void* o, void* value, void* key) {
	@autoreleasepool {
		[(NSObject*)o takeStoredValue:value forKey:key];
	}
}
BOOL
NSObject_inst_IsCaseInsensitiveLike(void* o, void* object) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSObject*)o isCaseInsensitiveLike:object];
	}
	return ret;

}
void* _Nullable
NSObject_inst_ValueWithName(void* o, void* name, void* key) {
	NSObject* _Nullable ret;
	@autoreleasepool {
		ret = [(NSObject*)o valueWithName:name inPropertyWithKey:key];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
BOOL
NSObject_inst_DoesContain(void* o, void* object) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSObject*)o doesContain:object];
	}
	return ret;

}
FourCharCode
NSObject_inst_ClassCode(void* o) {
	FourCharCode ret;
	@autoreleasepool {
		ret = [(NSObject*)o classCode];
	}
	return ret;

}
void* _Nullable
NSObject_inst_ScriptingValueForSpecifier(void* o, void* objectSpecifier) {
	NSObject* _Nullable ret;
	@autoreleasepool {
		ret = [(NSObject*)o scriptingValueForSpecifier:objectSpecifier];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
BOOL
NSObject_inst_IsGreaterThanOrEqualTo(void* o, void* object) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSObject*)o isGreaterThanOrEqualTo:object];
	}
	return ret;

}
BOOL
NSObject_inst_ScriptingIsEqualTo(void* o, void* object) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSObject*)o scriptingIsEqualTo:object];
	}
	return ret;

}
void* _Nullable
NSObject_inst_ObjectSpecifier(void* o) {
	NSScriptObjectSpecifier* _Nullable ret;
	@autoreleasepool {
		ret = [(NSObject*)o objectSpecifier];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void* _Nullable
NSObject_inst_ValueForKey(void* o, void* key) {
	NSObject* _Nullable ret;
	@autoreleasepool {
		ret = [(NSObject*)o valueForKey:key];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void*
NSObject_inst_Superclass(void* o) {
	Class ret;
	@autoreleasepool {
		ret = [(NSObject*)o superclass];
	}
	return ret;

}
void* _Nullable
NSObject_inst_ReplacementObjectForKeyedArchiver(void* o, void* archiver) {
	NSObject* _Nullable ret;
	@autoreleasepool {
		ret = [(NSObject*)o replacementObjectForKeyedArchiver:archiver];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void
NSObject_inst_HandleTakeValue(void* o, void* value, void* key) {
	@autoreleasepool {
		[(NSObject*)o handleTakeValue:value forUnboundKey:key];
	}
}
void
NSObject_inst_URLResourceDidFinishLoading(void* o, void* sender) {
	@autoreleasepool {
		[(NSObject*)o URLResourceDidFinishLoading:sender];
	}
}
void* _Nullable
NSObject_inst_InverseForRelationshipKey(void* o, void* relationshipKey) {
	NSString* _Nullable ret;
	@autoreleasepool {
		ret = [(NSObject*)o inverseForRelationshipKey:relationshipKey];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void
NSObject_inst_SetNilValueForKey(void* o, void* key) {
	@autoreleasepool {
		[(NSObject*)o setNilValueForKey:key];
	}
}
void*
NSObject_inst_Autorelease(void* o) {
	NSObject* ret;
	@autoreleasepool {
		ret = [(NSObject*)o autorelease];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void* _Nonnull
NSObject_inst_ToOneRelationshipKeys(void* o) {
	NSArray* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSObject*)o toOneRelationshipKeys];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
BOOL
NSObject_inst_IsLessThan(void* o, void* object) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSObject*)o isLessThan:object];
	}
	return ret;

}
void*
NSObject_inst_ForwardingTargetForSelector(void* o, void* aSelector) {
	NSObject* ret;
	@autoreleasepool {
		ret = [(NSObject*)o forwardingTargetForSelector:aSelector];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void
NSObject_inst_TakeValuesFromDictionary(void* o, void* properties) {
	@autoreleasepool {
		[(NSObject*)o takeValuesFromDictionary:properties];
	}
}
void
NSObject_inst_SetScriptingProperties(void* o, void* scriptingProperties) {
	@autoreleasepool {
		[(NSObject*)o setScriptingProperties:scriptingProperties];
	}
}
void* _Nullable
NSObject_inst_ObservationInfo(void* o) {
	void* _Nullable ret;
	@autoreleasepool {
		ret = [(NSObject*)o observationInfo];
	}
	return ret;

}
BOOL
NSObject_inst_IsNotEqualTo(void* o, void* object) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSObject*)o isNotEqualTo:object];
	}
	return ret;

}
void* _Nullable
NSObject_inst_HandleQueryWithUnboundKey(void* o, void* key) {
	NSObject* _Nullable ret;
	@autoreleasepool {
		ret = [(NSObject*)o handleQueryWithUnboundKey:key];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void* _Nullable
NSObject_inst_CopyScriptingValue(void* o, void* value, void* key, void* properties) {
	NSObject* _Nullable ret;
	@autoreleasepool {
		ret = [(NSObject*)o copyScriptingValue:value forKey:key withProperties:properties];
	}
	return ret;

}
void* _Nonnull
NSObject_inst_ValuesForKeys(void* o, void* keys) {
	NSDictionary* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSObject*)o valuesForKeys:keys];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void*
NSObject_inst_Init(void* o) {
	NSObject* ret;
	@autoreleasepool {
		ret = [(NSObject*)o init];
	}
	return ret;

}
void* _Nonnull
NSObject_inst_ClassName(void* o) {
	NSString* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSObject*)o className];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
BOOL
NSObject_inst_ScriptingIsLessThanOrEqualTo(void* o, void* object) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSObject*)o scriptingIsLessThanOrEqualTo:object];
	}
	return ret;

}
void* _Nullable
NSObject_inst_IndicesOfObjectsByEvaluatingObjectSpecifier(void* o, void* specifier) {
	NSArray* _Nullable ret;
	@autoreleasepool {
		ret = [(NSObject*)o indicesOfObjectsByEvaluatingObjectSpecifier:specifier];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void* _Nonnull
NSObject_inst_AutoContentAccessingProxy(void* o) {
	NSObject* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSObject*)o autoContentAccessingProxy];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
BOOL
NSObject_inst_ScriptingEndsWith(void* o, void* object) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSObject*)o scriptingEndsWith:object];
	}
	return ret;

}
BOOL
NSObject_inst_ScriptingIsGreaterThan(void* o, void* object) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSObject*)o scriptingIsGreaterThan:object];
	}
	return ret;

}
void* _Nonnull
NSObject_inst_MutableArrayValueForKey(void* o, void* key) {
	NSMutableArray* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSObject*)o mutableArrayValueForKey:key];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void*
NSObject_inst_Zone(void* o) {
	struct _NSZone* ret;
	@autoreleasepool {
		ret = [(NSObject*)o zone];
	}
	return ret;

}
void* _Nonnull
NSObject_inst_MutableArrayValueForKeyPath(void* o, void* keyPath) {
	NSMutableArray* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSObject*)o mutableArrayValueForKeyPath:keyPath];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void* _Nullable
NSObject_inst_ValueForUndefinedKey(void* o, void* key) {
	NSObject* _Nullable ret;
	@autoreleasepool {
		ret = [(NSObject*)o valueForUndefinedKey:key];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
BOOL
NSObject_inst_IsProxy(void* o) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSObject*)o isProxy];
	}
	return ret;

}
BOOL
NSObject_inst_ScriptingContains(void* o, void* object) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSObject*)o scriptingContains:object];
	}
	return ret;

}
void*
NSObject_inst_PerformSelector(void* o, void* aSelector) {
	NSObject* ret;
	@autoreleasepool {
		ret = [(NSObject*)o performSelector:aSelector];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void*
NSObject_inst_PerformSelectorWithObject(void* o, void* aSelector, void* object) {
	NSObject* ret;
	@autoreleasepool {
		ret = [(NSObject*)o performSelector:aSelector withObject:object];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void
NSObject_inst_PerformSelectorWithObjectAfterDelay(void* o, void* aSelector, void* anArgument, NSTimeInterval delay) {
	@autoreleasepool {
		[(NSObject*)o performSelector:aSelector withObject:anArgument afterDelay:delay];
	}
}
void*
NSObject_inst_PerformSelectorWithObjectWithObject(void* o, void* aSelector, void* object1, void* object2) {
	NSObject* ret;
	@autoreleasepool {
		ret = [(NSObject*)o performSelector:aSelector withObject:object1 withObject:object2];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void
NSObject_inst_PerformSelectorWithObjectAfterDelayInModes(void* o, void* aSelector, void* anArgument, NSTimeInterval delay, void* modes) {
	@autoreleasepool {
		[(NSObject*)o performSelector:aSelector withObject:anArgument afterDelay:delay inModes:modes];
	}
}
void
NSObject_inst_PerformSelectorOnThread(void* o, void* aSelector, void* thr, void* arg, BOOL wait) {
	@autoreleasepool {
		[(NSObject*)o performSelector:aSelector onThread:thr withObject:arg waitUntilDone:wait];
	}
}
void
NSObject_inst_PerformSelectorOnThreadWithObject(void* o, void* aSelector, void* thr, void* arg, BOOL wait, void* array) {
	@autoreleasepool {
		[(NSObject*)o performSelector:aSelector onThread:thr withObject:arg waitUntilDone:wait modes:array];
	}
}
void
NSObject_inst_RemoveValueAtIndex(void* o, NSUInteger index, void* key) {
	@autoreleasepool {
		[(NSObject*)o removeValueAtIndex:index fromPropertyWithKey:key];
	}
}
NSUInteger
NSObject_inst_Hash(void* o) {
	NSUInteger ret;
	@autoreleasepool {
		ret = [(NSObject*)o hash];
	}
	return ret;

}
void*
NSObject_inst_Class(void* o) {
	Class ret;
	@autoreleasepool {
		ret = [(NSObject*)o class];
	}
	return ret;

}
void*
NSObject_inst_DebugDescription(void* o) {
	NSString* ret;
	@autoreleasepool {
		ret = [(NSObject*)o debugDescription];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void* _Nullable
NSObject_inst_AwakeAfterUsingCoder(void* o, void* aDecoder) {
	NSObject* _Nullable ret;
	@autoreleasepool {
		ret = [(NSObject*)o awakeAfterUsingCoder:aDecoder];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void
NSObject_inst_DidChangeValueForKey(void* o, void* key) {
	@autoreleasepool {
		[(NSObject*)o didChangeValueForKey:key];
	}
}
void
NSObject_inst_DidChangeValueForKeyWithSetMutation(void* o, void* key, NSKeyValueSetMutationKind mutationKind, void* objects) {
	@autoreleasepool {
		[(NSObject*)o didChangeValueForKey:key withSetMutation:mutationKind usingObjects:objects];
	}
}
void* _Nonnull
NSObject_inst_MutableOrderedSetValueForKey(void* o, void* key) {
	NSMutableOrderedSet* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSObject*)o mutableOrderedSetValueForKey:key];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void
NSObject_inst_PerformSelectorInBackground(void* o, void* aSelector, void* arg) {
	@autoreleasepool {
		[(NSObject*)o performSelectorInBackground:aSelector withObject:arg];
	}
}
void* _Nonnull
NSObject_inst_ClassForCoder(void* o) {
	Class _Nonnull ret;
	@autoreleasepool {
		ret = [(NSObject*)o classForCoder];
	}
	return ret;

}
BOOL
NSObject_inst_ScriptingBeginsWith(void* o, void* object) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSObject*)o scriptingBeginsWith:object];
	}
	return ret;

}
void* _Nullable
NSObject_inst_CoerceValue(void* o, void* value, void* key) {
	NSObject* _Nullable ret;
	@autoreleasepool {
		ret = [(NSObject*)o coerceValue:value forKey:key];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void* _Nullable
NSObject_inst_ReplacementObjectForArchiver(void* o, void* archiver) {
	NSObject* _Nullable ret;
	@autoreleasepool {
		ret = [(NSObject*)o replacementObjectForArchiver:archiver];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
BOOL
NSObject_inst_IsEqual(void* o, void* object) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSObject*)o isEqual:object];
	}
	return ret;

}
BOOL
NSObject_inst_IsGreaterThan(void* o, void* object) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSObject*)o isGreaterThan:object];
	}
	return ret;

}
void* _Nonnull
NSObject_inst_AttributeKeys(void* o) {
	NSArray* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSObject*)o attributeKeys];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void* _Nullable
NSObject_inst_ReplacementObjectForCoder(void* o, void* aCoder) {
	NSObject* _Nullable ret;
	@autoreleasepool {
		ret = [(NSObject*)o replacementObjectForCoder:aCoder];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
BOOL
NSObject_inst_RespondsToSelector(void* o, void* aSelector) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSObject*)o respondsToSelector:aSelector];
	}
	return ret;

}
void* _Nullable
NSObject_inst_ValueWithUniqueID(void* o, void* uniqueID, void* key) {
	NSObject* _Nullable ret;
	@autoreleasepool {
		ret = [(NSObject*)o valueWithUniqueID:uniqueID inPropertyWithKey:key];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
BOOL
NSObject_inst_IsLessThanOrEqualTo(void* o, void* object) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSObject*)o isLessThanOrEqualTo:object];
	}
	return ret;

}
void
NSObject_inst_PerformSelectorOnMainThreadWithObject(void* o, void* aSelector, void* arg, BOOL wait) {
	@autoreleasepool {
		[(NSObject*)o performSelectorOnMainThread:aSelector withObject:arg waitUntilDone:wait];
	}
}
void
NSObject_inst_PerformSelectorOnMainThreadWithObjectWaitUntilDone(void* o, void* aSelector, void* arg, BOOL wait, void* array) {
	@autoreleasepool {
		[(NSObject*)o performSelectorOnMainThread:aSelector withObject:arg waitUntilDone:wait modes:array];
	}
}
void* _Nonnull
NSString_WritableTypeIdentifiersForItemProvider() {
	NSArray* _Nonnull ret;
	@autoreleasepool {
		ret = [NSString writableTypeIdentifiersForItemProvider];
	}
	return ret;

}
void* _Nonnull
NSString_StringWithCharacters(void* characters, NSUInteger length) {
	NSString* _Nonnull ret;
	@autoreleasepool {
		ret = [NSString stringWithCharacters:characters length:length];
		if(ret != nil) { [ret retain]; }
	}
	return ret;

}
void
NSString_SetVersion(NSInteger aVersion) {
	@autoreleasepool {
		[NSString setVersion:aVersion];
	}
}
void
NSString_Load() {
	@autoreleasepool {
		[NSString load];
	}
}
void
NSString_SetKeys(void* keys, void* dependentKey) {
	@autoreleasepool {
		[NSString setKeys:keys triggerChangeNotificationsForDependentKey:dependentKey];
	}
}
BOOL
NSString_SupportsSecureCoding() {
	BOOL ret;
	@autoreleasepool {
		ret = [NSString supportsSecureCoding];
	}
	return ret;

}
void* _Nonnull
NSString_StringWithFormat(void* format, void* object) {
	NSObject** arr = object;
	NSString* _Nonnull ret;
	@autoreleasepool {
		ret = [NSString stringWithFormat:format , arr[0], arr[1], arr[2], arr[3], arr[4], arr[5], arr[6], arr[7], arr[8], arr[9], arr[10], arr[11], arr[12], arr[13], arr[14], arr[15], nil];
		if(ret != nil) { [ret retain]; }
	}
	return ret;

}
BOOL
NSString_IsSubclassOfClass(void* aClass) {
	BOOL ret;
	@autoreleasepool {
		ret = [NSString isSubclassOfClass:aClass];
	}
	return ret;

}
NSInteger
NSString_Version() {
	NSInteger ret;
	@autoreleasepool {
		ret = [NSString version];
	}
	return ret;

}
void*
NSString_AllocWithZone(void* zone) {
	return [NSString allocWithZone:zone];
}
void*
NSString_Superclass() {
	Class ret;
	@autoreleasepool {
		ret = [NSString superclass];
	}
	return ret;

}
BOOL
NSString_AutomaticallyNotifiesObserversForKey(void* key) {
	BOOL ret;
	@autoreleasepool {
		ret = [NSString automaticallyNotifiesObserversForKey:key];
	}
	return ret;

}
void*
NSString_MutableCopyWithZone(void* zone) {
	NSObject* ret;
	@autoreleasepool {
		ret = [NSString mutableCopyWithZone:zone];
	}
	return ret;

}
void*
NSString_New() {
	NSString* ret;
	@autoreleasepool {
		ret = [NSString new];
	}
	return ret;

}
void
NSString_CancelPreviousPerformRequestsWithTarget(void* aTarget) {
	@autoreleasepool {
		[NSString cancelPreviousPerformRequestsWithTarget:aTarget];
	}
}
void
NSString_CancelPreviousPerformRequestsWithTargetSelector(void* aTarget, void* aSelector, void* anArgument) {
	@autoreleasepool {
		[NSString cancelPreviousPerformRequestsWithTarget:aTarget selector:aSelector object:anArgument];
	}
}
BOOL
NSString_InstancesRespondToSelector(void* aSelector) {
	BOOL ret;
	@autoreleasepool {
		ret = [NSString instancesRespondToSelector:aSelector];
	}
	return ret;

}
const void* _Nonnull
NSString_AvailableStringEncodings() {
	const NSStringEncoding* _Nonnull ret;
	@autoreleasepool {
		ret = [NSString availableStringEncodings];
	}
	return ret;

}
BOOL
NSString_AccessInstanceVariablesDirectly() {
	BOOL ret;
	@autoreleasepool {
		ret = [NSString accessInstanceVariablesDirectly];
	}
	return ret;

}
void* _Nonnull
NSString_KeyPathsForValuesAffectingValueForKey(void* key) {
	NSSet* _Nonnull ret;
	@autoreleasepool {
		ret = [NSString keyPathsForValuesAffectingValueForKey:key];
	}
	return ret;

}
void* _Nonnull
NSString_String() {
	NSString* _Nonnull ret;
	@autoreleasepool {
		ret = [NSString string];
		if(ret != nil) { [ret retain]; }
	}
	return ret;

}
void* _Nonnull
NSString_PathWithComponents(void* components) {
	NSString* _Nonnull ret;
	@autoreleasepool {
		ret = [NSString pathWithComponents:components];
	}
	return ret;

}
void* _Nonnull
NSString_LocalizedStringWithFormat(void* format, void* object) {
	NSObject** arr = object;
	NSString* _Nonnull ret;
	@autoreleasepool {
		ret = [NSString localizedStringWithFormat:format , arr[0], arr[1], arr[2], arr[3], arr[4], arr[5], arr[6], arr[7], arr[8], arr[9], arr[10], arr[11], arr[12], arr[13], arr[14], arr[15], nil];
	}
	return ret;

}
void* _Nullable
NSString_StringWithContentsOfFile(void* path) {
	NSObject* _Nullable ret;
	@autoreleasepool {
		ret = [NSString stringWithContentsOfFile:path];
		if(ret != nil) { [ret retain]; }
	}
	return ret;

}
void* _Nullable
NSString_StringWithContentsOfFileEncoding(void* path, NSStringEncoding enc, void** error) {
	NSString* _Nullable ret;
	@autoreleasepool {
		ret = [NSString stringWithContentsOfFile:path encoding:enc error:(NSError* _Nullable* _Nullable)error];
		if(ret != nil) { [ret retain]; }
		for(int i=0;i<1;i++) {
			if(error[i] == 0) { break; }
			[(id)error[i] retain];
		}
	
	}
	return ret;

}
void* _Nullable
NSString_StringWithContentsOfFileUsedEncoding(void* path, void* enc, void** error) {
	NSString* _Nullable ret;
	@autoreleasepool {
		ret = [NSString stringWithContentsOfFile:path usedEncoding:enc error:(NSError* _Nullable* _Nullable)error];
		if(ret != nil) { [ret retain]; }
		for(int i=0;i<1;i++) {
			if(error[i] == 0) { break; }
			[(id)error[i] retain];
		}
	
	}
	return ret;

}
NSItemProviderRepresentationVisibility
NSString_ItemProviderVisibilityForRepresentationWithTypeIdentifier(void* typeIdentifier) {
	NSItemProviderRepresentationVisibility ret;
	@autoreleasepool {
		ret = [NSString itemProviderVisibilityForRepresentationWithTypeIdentifier:typeIdentifier];
	}
	return ret;

}
void*
NSString_Class() {
	Class ret;
	@autoreleasepool {
		ret = [NSString class];
	}
	return ret;

}
NSUInteger
NSString_Hash() {
	NSUInteger ret;
	@autoreleasepool {
		ret = [NSString hash];
	}
	return ret;

}
NSStringEncoding
NSString_DefaultCStringEncoding() {
	NSStringEncoding ret;
	@autoreleasepool {
		ret = [NSString defaultCStringEncoding];
	}
	return ret;

}
void* _Nonnull
NSString_StringWithString(void* string) {
	NSString* _Nonnull ret;
	@autoreleasepool {
		ret = [NSString stringWithString:string];
		if(ret != nil) { [ret retain]; }
	}
	return ret;

}
BOOL
NSString_UseStoredAccessor() {
	BOOL ret;
	@autoreleasepool {
		ret = [NSString useStoredAccessor];
	}
	return ret;

}
BOOL
NSString_ConformsToProtocol(void* protocol) {
	BOOL ret;
	@autoreleasepool {
		ret = [NSString conformsToProtocol:protocol];
	}
	return ret;

}
NSStringEncoding
NSString_StringEncodingForData(void* data, void* opts, void** string, void* usedLossyConversion) {
	NSStringEncoding ret;
	@autoreleasepool {
		ret = [NSString stringEncodingForData:data encodingOptions:opts convertedString:(NSString* _Nullable* _Nullable)string usedLossyConversion:usedLossyConversion];
		for(int i=0;i<1;i++) {
			if(string[i] == 0) { break; }
			[(id)string[i] retain];
		}
	
	}
	return ret;

}
void* _Nonnull
NSString_LocalizedNameOfStringEncoding(NSStringEncoding encoding) {
	NSString* _Nonnull ret;
	@autoreleasepool {
		ret = [NSString localizedNameOfStringEncoding:encoding];
	}
	return ret;

}
void* _Nonnull
NSString_ClassFallbacksForKeyedArchiver() {
	NSArray* _Nonnull ret;
	@autoreleasepool {
		ret = [NSString classFallbacksForKeyedArchiver];
	}
	return ret;

}
BOOL
NSString_ResolveInstanceMethod(void* sel) {
	BOOL ret;
	@autoreleasepool {
		ret = [NSString resolveInstanceMethod:sel];
	}
	return ret;

}
void*
NSString_Description() {
	NSString* ret;
	@autoreleasepool {
		ret = [NSString description];
	}
	return ret;

}
void*
NSString_CopyWithZone(void* zone) {
	NSObject* ret;
	@autoreleasepool {
		ret = [NSString copyWithZone:zone];
	}
	return ret;

}
void* _Nonnull
NSString_ReadableTypeIdentifiersForItemProvider() {
	NSArray* _Nonnull ret;
	@autoreleasepool {
		ret = [NSString readableTypeIdentifiersForItemProvider];
	}
	return ret;

}
void* _Nullable
NSString_StringWithUTF8String(void* nullTerminatedCString) {
	NSString* _Nullable ret;
	@autoreleasepool {
		ret = [NSString stringWithUTF8String:nullTerminatedCString];
		if(ret != nil) { [ret retain]; }
	}
	return ret;

}
void*
NSString_InstanceMethodSignatureForSelector(void* aSelector) {
	NSMethodSignature* ret;
	@autoreleasepool {
		ret = [NSString instanceMethodSignatureForSelector:aSelector];
	}
	return ret;

}
void* _Nonnull
NSString_ClassForKeyedUnarchiver() {
	Class _Nonnull ret;
	@autoreleasepool {
		ret = [NSString classForKeyedUnarchiver];
	}
	return ret;

}
BOOL
NSString_ResolveClassMethod(void* sel) {
	BOOL ret;
	@autoreleasepool {
		ret = [NSString resolveClassMethod:sel];
	}
	return ret;

}
void* _Nullable
NSString_ObjectWithItemProviderData(void* data, void* typeIdentifier, void** outError) {
	NSString* _Nullable ret;
	@autoreleasepool {
		ret = [NSString objectWithItemProviderData:data typeIdentifier:typeIdentifier error:(NSError* _Nullable* _Nullable)outError];
		for(int i=0;i<1;i++) {
			if(outError[i] == 0) { break; }
			[(id)outError[i] retain];
		}
	
	}
	return ret;

}
void* _Nullable
NSString_StringWithCString(void* bytes) {
	NSObject* _Nullable ret;
	@autoreleasepool {
		ret = [NSString stringWithCString:bytes];
		if(ret != nil) { [ret retain]; }
	}
	return ret;

}
void* _Nullable
NSString_StringWithCStringEncoding(void* cString, NSStringEncoding enc) {
	NSString* _Nullable ret;
	@autoreleasepool {
		ret = [NSString stringWithCString:cString encoding:enc];
		if(ret != nil) { [ret retain]; }
	}
	return ret;

}
void* _Nullable
NSString_StringWithCStringLength(void* bytes, NSUInteger length) {
	NSObject* _Nullable ret;
	@autoreleasepool {
		ret = [NSString stringWithCString:bytes length:length];
		if(ret != nil) { [ret retain]; }
	}
	return ret;

}
void* _Nullable
NSString_StringWithContentsOfURL(void* url) {
	NSObject* _Nullable ret;
	@autoreleasepool {
		ret = [NSString stringWithContentsOfURL:url];
		if(ret != nil) { [ret retain]; }
	}
	return ret;

}
void* _Nullable
NSString_StringWithContentsOfURLEncoding(void* url, NSStringEncoding enc, void** error) {
	NSString* _Nullable ret;
	@autoreleasepool {
		ret = [NSString stringWithContentsOfURL:url encoding:enc error:(NSError* _Nullable* _Nullable)error];
		if(ret != nil) { [ret retain]; }
		for(int i=0;i<1;i++) {
			if(error[i] == 0) { break; }
			[(id)error[i] retain];
		}
	
	}
	return ret;

}
void* _Nullable
NSString_StringWithContentsOfURLUsedEncoding(void* url, void* enc, void** error) {
	NSString* _Nullable ret;
	@autoreleasepool {
		ret = [NSString stringWithContentsOfURL:url usedEncoding:enc error:(NSError* _Nullable* _Nullable)error];
		if(ret != nil) { [ret retain]; }
		for(int i=0;i<1;i++) {
			if(error[i] == 0) { break; }
			[(id)error[i] retain];
		}
	
	}
	return ret;

}
void*
NSString_Alloc() {
	return [NSString alloc];
}
void*
NSString_DebugDescription() {
	NSString* ret;
	@autoreleasepool {
		ret = [NSString debugDescription];
	}
	return ret;

}
void* _Nullable
NSString_inst_InitWithBytesNoCopy(void* o, void* bytes, NSUInteger len, NSStringEncoding encoding, BOOL freeBuffer) {
	NSString* _Nullable ret;
	@autoreleasepool {
		ret = [(NSString*)o initWithBytesNoCopy:bytes length:len encoding:encoding freeWhenDone:freeBuffer];
	}
	return ret;

}
NSStringEncoding
NSString_inst_FastestEncoding(void* o) {
	NSStringEncoding ret;
	@autoreleasepool {
		ret = [(NSString*)o fastestEncoding];
	}
	return ret;

}
void
NSString_inst_GetCString(void* o, void* bytes) {
	@autoreleasepool {
		[(NSString*)o getCString:bytes];
	}
}
void
NSString_inst_GetCStringMaxLength(void* o, void* bytes, NSUInteger maxLength) {
	@autoreleasepool {
		[(NSString*)o getCString:bytes maxLength:maxLength];
	}
}
BOOL
NSString_inst_GetCStringMaxLengthEncoding(void* o, void* buffer, NSUInteger maxBufferCount, NSStringEncoding encoding) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSString*)o getCString:buffer maxLength:maxBufferCount encoding:encoding];
	}
	return ret;

}
void
NSString_inst_GetCStringMaxLengthRange(void* o, void* bytes, NSUInteger maxLength, NSRange aRange, void* leftoverRange) {
	@autoreleasepool {
		[(NSString*)o getCString:bytes maxLength:maxLength range:aRange remainingRange:leftoverRange];
	}
}
void* _Nullable
NSString_inst_ValueWithName(void* o, void* name, void* key) {
	NSObject* _Nullable ret;
	@autoreleasepool {
		ret = [(NSString*)o valueWithName:name inPropertyWithKey:key];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void* _Nonnull
NSString_inst_VariantFittingPresentationWidth(void* o, NSInteger width) {
	NSString* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o variantFittingPresentationWidth:width];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
BOOL
NSString_inst_ScriptingIsLessThan(void* o, void* object) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSString*)o scriptingIsLessThan:object];
	}
	return ret;

}
void* _Nullable
NSString_inst_ValueAtIndex(void* o, NSUInteger index, void* key) {
	NSObject* _Nullable ret;
	@autoreleasepool {
		ret = [(NSString*)o valueAtIndex:index inPropertyWithKey:key];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void* _Nullable
NSString_inst_IndicesOfObjectsByEvaluatingObjectSpecifier(void* o, void* specifier) {
	NSArray* _Nullable ret;
	@autoreleasepool {
		ret = [(NSString*)o indicesOfObjectsByEvaluatingObjectSpecifier:specifier];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void* _Nonnull
NSString_inst_MutableArrayValueForKey(void* o, void* key) {
	NSMutableArray* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o mutableArrayValueForKey:key];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void*
NSString_inst_PerformSelector(void* o, void* aSelector) {
	NSObject* ret;
	@autoreleasepool {
		ret = [(NSString*)o performSelector:aSelector];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void*
NSString_inst_PerformSelectorWithObject(void* o, void* aSelector, void* object) {
	NSObject* ret;
	@autoreleasepool {
		ret = [(NSString*)o performSelector:aSelector withObject:object];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void
NSString_inst_PerformSelectorWithObjectAfterDelay(void* o, void* aSelector, void* anArgument, NSTimeInterval delay) {
	@autoreleasepool {
		[(NSString*)o performSelector:aSelector withObject:anArgument afterDelay:delay];
	}
}
void*
NSString_inst_PerformSelectorWithObjectWithObject(void* o, void* aSelector, void* object1, void* object2) {
	NSObject* ret;
	@autoreleasepool {
		ret = [(NSString*)o performSelector:aSelector withObject:object1 withObject:object2];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void
NSString_inst_PerformSelectorWithObjectAfterDelayInModes(void* o, void* aSelector, void* anArgument, NSTimeInterval delay, void* modes) {
	@autoreleasepool {
		[(NSString*)o performSelector:aSelector withObject:anArgument afterDelay:delay inModes:modes];
	}
}
void
NSString_inst_PerformSelectorOnThread(void* o, void* aSelector, void* thr, void* arg, BOOL wait) {
	@autoreleasepool {
		[(NSString*)o performSelector:aSelector onThread:thr withObject:arg waitUntilDone:wait];
	}
}
void
NSString_inst_PerformSelectorOnThreadWithObject(void* o, void* aSelector, void* thr, void* arg, BOOL wait, void* array) {
	@autoreleasepool {
		[(NSString*)o performSelector:aSelector onThread:thr withObject:arg waitUntilDone:wait modes:array];
	}
}
void* _Nonnull
NSString_inst_ClassForCoder(void* o) {
	Class _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o classForCoder];
	}
	return ret;

}
double
NSString_inst_DoubleValue(void* o) {
	double ret;
	@autoreleasepool {
		ret = [(NSString*)o doubleValue];
	}
	return ret;

}
float
NSString_inst_FloatValue(void* o) {
	float ret;
	@autoreleasepool {
		ret = [(NSString*)o floatValue];
	}
	return ret;

}
void* _Nonnull
NSString_inst_ToManyRelationshipKeys(void* o) {
	NSArray* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o toManyRelationshipKeys];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void* _Nonnull
NSString_inst_MutableOrderedSetValueForKeyPath(void* o, void* keyPath) {
	NSMutableOrderedSet* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o mutableOrderedSetValueForKeyPath:keyPath];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void*
NSString_inst_ForwardingTargetForSelector(void* o, void* aSelector) {
	NSObject* ret;
	@autoreleasepool {
		ret = [(NSString*)o forwardingTargetForSelector:aSelector];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void*
NSString_inst_Zone(void* o) {
	struct _NSZone* ret;
	@autoreleasepool {
		ret = [(NSString*)o zone];
	}
	return ret;

}
void* _Nonnull
NSString_inst_StringByDeletingPathExtension(void* o) {
	NSString* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o stringByDeletingPathExtension];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void* _Nonnull
NSString_inst_LinguisticTagsInRange(void* o, NSRange range, void* scheme, NSLinguisticTaggerOptions options, void* orthography, void** tokenRanges) {
	NSArray* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o linguisticTagsInRange:range scheme:scheme options:options orthography:orthography tokenRanges:(NSArray <NSValue*>* _Nullable* _Nullable)tokenRanges];
		if (ret != nil && ret != o) { [ret retain]; }
		for(int i=0;i<1;i++) {
			if(tokenRanges[i] == 0) { break; }
			[(id)tokenRanges[i] retain];
		}
	
	}
	return ret;

}
void* _Nullable
NSString_inst_InitWithCoder(void* o, void* aDecoder) {
	NSString* _Nullable ret;
	@autoreleasepool {
		ret = [(NSString*)o initWithCoder:aDecoder];
	}
	return ret;

}
NSUInteger
NSString_inst_CStringLength(void* o) {
	NSUInteger ret;
	@autoreleasepool {
		ret = [(NSString*)o cStringLength];
	}
	return ret;

}
void* _Nonnull
NSString_inst_StringByAppendingFormat(void* o, void* format, void* object) {
	NSObject** arr = object;
	NSString* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o stringByAppendingFormat:format , arr[0], arr[1], arr[2], arr[3], arr[4], arr[5], arr[6], arr[7], arr[8], arr[9], arr[10], arr[11], arr[12], arr[13], arr[14], arr[15], nil];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
BOOL
NSString_inst_LocalizedStandardContainsString(void* o, void* str) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSString*)o localizedStandardContainsString:str];
	}
	return ret;

}
void* _Nonnull
NSString_inst_SubstringToIndex(void* o, NSUInteger to) {
	NSString* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o substringToIndex:to];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void
NSString_inst_SetValuesForKeysWithDictionary(void* o, void* keyedValues) {
	@autoreleasepool {
		[(NSString*)o setValuesForKeysWithDictionary:keyedValues];
	}
}
void* _Nullable
NSString_inst_ValueForKeyPath(void* o, void* keyPath) {
	NSObject* _Nullable ret;
	@autoreleasepool {
		ret = [(NSString*)o valueForKeyPath:keyPath];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void
NSString_inst_TakeValueForKey(void* o, void* value, void* key) {
	@autoreleasepool {
		[(NSString*)o takeValue:value forKey:key];
	}
}
void
NSString_inst_TakeValueForKeyPath(void* o, void* value, void* keyPath) {
	@autoreleasepool {
		[(NSString*)o takeValue:value forKeyPath:keyPath];
	}
}
void* _Nonnull
NSString_inst_PrecomposedStringWithCanonicalMapping(void* o) {
	NSString* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o precomposedStringWithCanonicalMapping];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void* _Nonnull
NSString_inst_ComponentsSeparatedByCharactersInSet(void* o, void* separator) {
	NSArray* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o componentsSeparatedByCharactersInSet:separator];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void* _Nullable
NSString_inst_DataUsingEncoding(void* o, NSStringEncoding encoding) {
	NSData* _Nullable ret;
	@autoreleasepool {
		ret = [(NSString*)o dataUsingEncoding:encoding];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void* _Nullable
NSString_inst_DataUsingEncodingAllowLossyConversion(void* o, NSStringEncoding encoding, BOOL lossy) {
	NSData* _Nullable ret;
	@autoreleasepool {
		ret = [(NSString*)o dataUsingEncoding:encoding allowLossyConversion:lossy];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
BOOL
NSString_inst_IsEqual(void* o, void* object) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSString*)o isEqual:object];
	}
	return ret;

}
BOOL
NSString_inst_GetBytes(void* o, void* buffer, NSUInteger maxBufferCount, void* usedBufferCount, NSStringEncoding encoding, NSStringEncodingConversionOptions options, NSRange range, void* leftover) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSString*)o getBytes:buffer maxLength:maxBufferCount usedLength:usedBufferCount encoding:encoding options:options range:range remainingRange:leftover];
	}
	return ret;

}
NSUInteger
NSString_inst_CompletePathIntoString(void* o, void** outputName, BOOL flag, void** outputArray, void* filterTypes) {
	NSUInteger ret;
	@autoreleasepool {
		ret = [(NSString*)o completePathIntoString:(NSString* _Nullable* _Nullable)outputName caseSensitive:flag matchesIntoArray:(NSArray <NSString*>* _Nullable* _Nullable)outputArray filterTypes:filterTypes];
		for(int i=0;i<1;i++) {
			if(outputName[i] == 0) { break; }
			[(id)outputName[i] retain];
		}
	
	
		for(int i=0;i<1;i++) {
			if(outputArray[i] == 0) { break; }
			[(id)outputArray[i] retain];
		}
	
	}
	return ret;

}
void* _Nonnull
NSString_inst_ClassForPortCoder(void* o) {
	Class _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o classForPortCoder];
	}
	return ret;

}
BOOL
NSString_inst_IsMemberOfClass(void* o, void* aClass) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSString*)o isMemberOfClass:aClass];
	}
	return ret;

}
void* _Nonnull
NSString_inst_StringByAppendingPathComponent(void* o, void* str) {
	NSString* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o stringByAppendingPathComponent:str];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
BOOL
NSString_inst_ScriptingEndsWith(void* o, void* object) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSString*)o scriptingEndsWith:object];
	}
	return ret;

}
void* _Nullable
NSString_inst_ReplacementObjectForArchiver(void* o, void* archiver) {
	NSObject* _Nullable ret;
	@autoreleasepool {
		ret = [(NSString*)o replacementObjectForArchiver:archiver];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
BOOL
NSString_inst_BoolValue(void* o) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSString*)o boolValue];
	}
	return ret;

}
NSItemProviderRepresentationVisibility
NSString_inst_ItemProviderVisibilityForRepresentationWithTypeIdentifier(void* o, void* typeIdentifier) {
	NSItemProviderRepresentationVisibility ret;
	@autoreleasepool {
		ret = [(NSString*)o itemProviderVisibilityForRepresentationWithTypeIdentifier:typeIdentifier];
	}
	return ret;

}
BOOL
NSString_inst_LocalizedCaseInsensitiveContainsString(void* o, void* str) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSString*)o localizedCaseInsensitiveContainsString:str];
	}
	return ret;

}
void* _Nullable
NSString_inst_HandleQueryWithUnboundKey(void* o, void* key) {
	NSObject* _Nullable ret;
	@autoreleasepool {
		ret = [(NSString*)o handleQueryWithUnboundKey:key];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
BOOL
NSString_inst_ScriptingIsGreaterThan(void* o, void* object) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSString*)o scriptingIsGreaterThan:object];
	}
	return ret;

}
void
NSString_inst_RemoveValueAtIndex(void* o, NSUInteger index, void* key) {
	@autoreleasepool {
		[(NSString*)o removeValueAtIndex:index fromPropertyWithKey:key];
	}
}
NSUInteger
NSString_inst_MaximumLengthOfBytesUsingEncoding(void* o, NSStringEncoding enc) {
	NSUInteger ret;
	@autoreleasepool {
		ret = [(NSString*)o maximumLengthOfBytesUsingEncoding:enc];
	}
	return ret;

}
void* _Nonnull
NSString_inst_CapitalizedStringWithLocale(void* o, void* locale) {
	NSString* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o capitalizedStringWithLocale:locale];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void* _Nonnull
NSString_inst_UppercaseStringWithLocale(void* o, void* locale) {
	NSString* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o uppercaseStringWithLocale:locale];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void* _Nonnull
NSString_inst_Description(void* o) {
	NSString* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o description];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
NSInteger
NSString_inst_IntegerValue(void* o) {
	NSInteger ret;
	@autoreleasepool {
		ret = [(NSString*)o integerValue];
	}
	return ret;

}
void* _Nullable
NSString_inst_ReplacementObjectForPortCoder(void* o, void* coder) {
	NSObject* _Nullable ret;
	@autoreleasepool {
		ret = [(NSString*)o replacementObjectForPortCoder:coder];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void
NSString_inst_DoesNotRecognizeSelector(void* o, void* aSelector) {
	@autoreleasepool {
		[(NSString*)o doesNotRecognizeSelector:aSelector];
	}
}
void* _Nonnull
NSString_inst_AutoContentAccessingProxy(void* o) {
	NSObject* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o autoContentAccessingProxy];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void* _Nullable
NSString_inst_InitWithUTF8String(void* o, void* nullTerminatedCString) {
	NSString* _Nullable ret;
	@autoreleasepool {
		ret = [(NSString*)o initWithUTF8String:nullTerminatedCString];
	}
	return ret;

}
void* _Nullable
NSString_inst_StringByAddingPercentEscapesUsingEncoding(void* o, NSStringEncoding enc) {
	NSString* _Nullable ret;
	@autoreleasepool {
		ret = [(NSString*)o stringByAddingPercentEscapesUsingEncoding:enc];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void* _Nonnull
NSString_inst_ClassName(void* o) {
	NSString* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o className];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void* _Nonnull
NSString_inst_ComponentsSeparatedByString(void* o, void* separator) {
	NSArray* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o componentsSeparatedByString:separator];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
BOOL
NSString_inst_ConformsToProtocol(void* o, void* aProtocol) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSString*)o conformsToProtocol:aProtocol];
	}
	return ret;

}
void* _Nullable
NSString_inst_ValueForKey(void* o, void* key) {
	NSObject* _Nullable ret;
	@autoreleasepool {
		ret = [(NSString*)o valueForKey:key];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void* _Nonnull
NSString_inst_StringByDeletingLastPathComponent(void* o) {
	NSString* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o stringByDeletingLastPathComponent];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void* _Nonnull
NSString_inst_StringByTrimmingCharactersInSet(void* o, void* set) {
	NSString* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o stringByTrimmingCharactersInSet:set];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void
NSString_inst_UnableToSetNilForKey(void* o, void* key) {
	@autoreleasepool {
		[(NSString*)o unableToSetNilForKey:key];
	}
}
void* _Nonnull
NSString_inst_DictionaryWithValuesForKeys(void* o, void* keys) {
	NSDictionary* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o dictionaryWithValuesForKeys:keys];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void*
NSString_inst_MutableCopy(void* o) {
	NSObject* ret;
	@autoreleasepool {
		ret = [(NSString*)o mutableCopy];
	}
	return ret;

}
NSRange
NSString_inst_RangeOfString(void* o, void* searchString) {
	NSRange ret;
	@autoreleasepool {
		ret = [(NSString*)o rangeOfString:searchString];
	}
	return ret;

}
NSRange
NSString_inst_RangeOfStringOptions(void* o, void* searchString, NSStringCompareOptions mask) {
	NSRange ret;
	@autoreleasepool {
		ret = [(NSString*)o rangeOfString:searchString options:mask];
	}
	return ret;

}
NSRange
NSString_inst_RangeOfStringOptionsRange(void* o, void* searchString, NSStringCompareOptions mask, NSRange rangeOfReceiverToSearch) {
	NSRange ret;
	@autoreleasepool {
		ret = [(NSString*)o rangeOfString:searchString options:mask range:rangeOfReceiverToSearch];
	}
	return ret;

}
NSRange
NSString_inst_RangeOfStringOptionsRangeLocale(void* o, void* searchString, NSStringCompareOptions mask, NSRange rangeOfReceiverToSearch, void* locale) {
	NSRange ret;
	@autoreleasepool {
		ret = [(NSString*)o rangeOfString:searchString options:mask range:rangeOfReceiverToSearch locale:locale];
	}
	return ret;

}
void* _Nonnull
NSString_inst_LowercaseString(void* o) {
	NSString* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o lowercaseString];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void*
NSString_inst_Retain(void* o) {
	NSString* ret;
	@autoreleasepool {
		ret = [(NSString*)o retain];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void
NSString_inst_ForwardInvocation(void* o, void* anInvocation) {
	@autoreleasepool {
		[(NSString*)o forwardInvocation:anInvocation];
	}
}
void* _Nullable
NSString_inst_ScriptingValueForSpecifier(void* o, void* objectSpecifier) {
	NSObject* _Nullable ret;
	@autoreleasepool {
		ret = [(NSString*)o scriptingValueForSpecifier:objectSpecifier];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void* _Nonnull
NSString_inst_PropertyList(void* o) {
	NSObject* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o propertyList];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
NSRange
NSString_inst_LocalizedStandardRangeOfString(void* o, void* str) {
	NSRange ret;
	@autoreleasepool {
		ret = [(NSString*)o localizedStandardRangeOfString:str];
	}
	return ret;

}
void*
NSString_inst_MethodSignatureForSelector(void* o, void* aSelector) {
	NSMethodSignature* ret;
	@autoreleasepool {
		ret = [(NSString*)o methodSignatureForSelector:aSelector];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void
NSString_inst_AddObserver(void* o, void* observer, void* keyPath, NSKeyValueObservingOptions options, void* context) {
	@autoreleasepool {
		[(NSString*)o addObserver:observer forKeyPath:keyPath options:options context:context];
	}
}
BOOL
NSString_inst_GetFileSystemRepresentation(void* o, void* cname, NSUInteger max) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSString*)o getFileSystemRepresentation:cname maxLength:max];
	}
	return ret;

}
void* _Nonnull
NSString_inst_StringByReplacingOccurrencesOfStringWithString(void* o, void* target, void* replacement) {
	NSString* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o stringByReplacingOccurrencesOfString:target withString:replacement];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void* _Nonnull
NSString_inst_StringByReplacingOccurrencesOfStringWithStringOptions(void* o, void* target, void* replacement, NSStringCompareOptions options, NSRange searchRange) {
	NSString* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o stringByReplacingOccurrencesOfString:target withString:replacement options:options range:searchRange];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
const void* _Nullable
NSString_inst_UTF8String(void* o) {
	const char* _Nullable ret;
	@autoreleasepool {
		ret = strdup([(NSString*)o UTF8String]);
	}
	return ret;

}
BOOL
NSString_inst_FileManagerShouldProceedAfterError(void* o, void* fm, void* errorInfo) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSString*)o fileManager:fm shouldProceedAfterError:errorInfo];
	}
	return ret;

}
void
NSString_inst_FileManagerWillProcessPath(void* o, void* fm, void* path) {
	@autoreleasepool {
		[(NSString*)o fileManager:fm willProcessPath:path];
	}
}
void* _Nonnull
NSString_inst_LowercaseStringWithLocale(void* o, void* locale) {
	NSString* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o lowercaseStringWithLocale:locale];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
NSRange
NSString_inst_RangeOfComposedCharacterSequencesForRange(void* o, NSRange range) {
	NSRange ret;
	@autoreleasepool {
		ret = [(NSString*)o rangeOfComposedCharacterSequencesForRange:range];
	}
	return ret;

}
BOOL
NSString_inst_ScriptingIsGreaterThanOrEqualTo(void* o, void* object) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSString*)o scriptingIsGreaterThanOrEqualTo:object];
	}
	return ret;

}
void* _Nonnull
NSString_inst_MutableOrderedSetValueForKey(void* o, void* key) {
	NSMutableOrderedSet* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o mutableOrderedSetValueForKey:key];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void
NSString_inst_PerformSelectorOnMainThreadWithObject(void* o, void* aSelector, void* arg, BOOL wait) {
	@autoreleasepool {
		[(NSString*)o performSelectorOnMainThread:aSelector withObject:arg waitUntilDone:wait];
	}
}
void
NSString_inst_PerformSelectorOnMainThreadWithObjectWaitUntilDone(void* o, void* aSelector, void* arg, BOOL wait, void* array) {
	@autoreleasepool {
		[(NSString*)o performSelectorOnMainThread:aSelector withObject:arg waitUntilDone:wait modes:array];
	}
}
void* _Nonnull
NSString_inst_DecomposedStringWithCanonicalMapping(void* o) {
	NSString* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o decomposedStringWithCanonicalMapping];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
NSRange
NSString_inst_LineRangeForRange(void* o, NSRange range) {
	NSRange ret;
	@autoreleasepool {
		ret = [(NSString*)o lineRangeForRange:range];
	}
	return ret;

}
unichar
NSString_inst_CharacterAtIndex(void* o, NSUInteger index) {
	unichar ret;
	@autoreleasepool {
		ret = [(NSString*)o characterAtIndex:index];
	}
	return ret;

}
void* _Nonnull
NSString_inst_ClassDescription(void* o) {
	NSClassDescription* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o classDescription];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void* _Nullable
NSString_inst_NewScriptingObjectOfClass(void* o, void* objectClass, void* key, void* contentsValue, void* properties) {
	NSObject* _Nullable ret;
	@autoreleasepool {
		ret = [(NSString*)o newScriptingObjectOfClass:objectClass forValueForKey:key withContentsValue:contentsValue properties:properties];
	}
	return ret;

}
void
NSString_inst_URLResourceDidCancelLoading(void* o, void* sender) {
	@autoreleasepool {
		[(NSString*)o URLResourceDidCancelLoading:sender];
	}
}
NSComparisonResult
NSString_inst_LocalizedStandardCompare(void* o, void* string) {
	NSComparisonResult ret;
	@autoreleasepool {
		ret = [(NSString*)o localizedStandardCompare:string];
	}
	return ret;

}
void* _Nullable
NSString_inst_StringByReplacingPercentEscapesUsingEncoding(void* o, NSStringEncoding enc) {
	NSString* _Nullable ret;
	@autoreleasepool {
		ret = [(NSString*)o stringByReplacingPercentEscapesUsingEncoding:enc];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void
NSString_inst_WillChangeValueForKey(void* o, void* key) {
	@autoreleasepool {
		[(NSString*)o willChangeValueForKey:key];
	}
}
void
NSString_inst_WillChangeValueForKeyWithSetMutation(void* o, void* key, NSKeyValueSetMutationKind mutationKind, void* objects) {
	@autoreleasepool {
		[(NSString*)o willChangeValueForKey:key withSetMutation:mutationKind usingObjects:objects];
	}
}
BOOL
NSString_inst_AttemptRecoveryFromErrorOptionIndex(void* o, void* error, NSUInteger recoveryOptionIndex) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSString*)o attemptRecoveryFromError:error optionIndex:recoveryOptionIndex];
	}
	return ret;

}
void
NSString_inst_AttemptRecoveryFromErrorOptionIndexDelegate(void* o, void* error, NSUInteger recoveryOptionIndex, void* delegate, void* didRecoverSelector, void* contextInfo) {
	@autoreleasepool {
		[(NSString*)o attemptRecoveryFromError:error optionIndex:recoveryOptionIndex delegate:delegate didRecoverSelector:didRecoverSelector contextInfo:contextInfo];
	}
}
FourCharCode
NSString_inst_ClassCode(void* o) {
	FourCharCode ret;
	@autoreleasepool {
		ret = [(NSString*)o classCode];
	}
	return ret;

}
NSComparisonResult
NSString_inst_Compare(void* o, void* string) {
	NSComparisonResult ret;
	@autoreleasepool {
		ret = [(NSString*)o compare:string];
	}
	return ret;

}
NSComparisonResult
NSString_inst_CompareOptions(void* o, void* string, NSStringCompareOptions mask) {
	NSComparisonResult ret;
	@autoreleasepool {
		ret = [(NSString*)o compare:string options:mask];
	}
	return ret;

}
NSComparisonResult
NSString_inst_CompareOptionsRange(void* o, void* string, NSStringCompareOptions mask, NSRange rangeOfReceiverToCompare) {
	NSComparisonResult ret;
	@autoreleasepool {
		ret = [(NSString*)o compare:string options:mask range:rangeOfReceiverToCompare];
	}
	return ret;

}
NSComparisonResult
NSString_inst_CompareOptionsRangeLocale(void* o, void* string, NSStringCompareOptions mask, NSRange rangeOfReceiverToCompare, void* locale) {
	NSComparisonResult ret;
	@autoreleasepool {
		ret = [(NSString*)o compare:string options:mask range:rangeOfReceiverToCompare locale:locale];
	}
	return ret;

}
BOOL
NSString_inst_HasSuffix(void* o, void* str) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSString*)o hasSuffix:str];
	}
	return ret;

}
void* _Nonnull
NSString_inst_StringByResolvingSymlinksInPath(void* o) {
	NSString* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o stringByResolvingSymlinksInPath];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void* _Nonnull
NSString_inst_InitWithFormat(void* o, void* format, void* object) {
	NSObject** arr = object;
	NSString* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o initWithFormat:format , arr[0], arr[1], arr[2], arr[3], arr[4], arr[5], arr[6], arr[7], arr[8], arr[9], arr[10], arr[11], arr[12], arr[13], arr[14], arr[15], nil];
	}
	return ret;

}
void* _Nonnull
NSString_inst_InitWithFormatLocale(void* o, void* format, void* locale, void* object) {
	NSObject** arr = object;
	NSString* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o initWithFormat:format locale:locale , arr[0], arr[1], arr[2], arr[3], arr[4], arr[5], arr[6], arr[7], arr[8], arr[9], arr[10], arr[11], arr[12], arr[13], arr[14], arr[15], nil];
	}
	return ret;

}
long long
NSString_inst_LongLongValue(void* o) {
	long long ret;
	@autoreleasepool {
		ret = [(NSString*)o longLongValue];
	}
	return ret;

}
void* _Nonnull
NSString_inst_StringByExpandingTildeInPath(void* o) {
	NSString* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o stringByExpandingTildeInPath];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void*
NSString_inst_Superclass(void* o) {
	Class ret;
	@autoreleasepool {
		ret = [(NSString*)o superclass];
	}
	return ret;

}
void
NSString_inst_PerformSelectorInBackground(void* o, void* aSelector, void* arg) {
	@autoreleasepool {
		[(NSString*)o performSelectorInBackground:aSelector withObject:arg];
	}
}
void* _Nonnull
NSString_inst_UppercaseString(void* o) {
	NSString* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o uppercaseString];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
BOOL
NSString_inst_IsEqualTo(void* o, void* object) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSString*)o isEqualTo:object];
	}
	return ret;

}
NSUInteger
NSString_inst_RetainCount(void* o) {
	NSUInteger ret;
	@autoreleasepool {
		ret = [(NSString*)o retainCount];
	}
	return ret;

}
void* _Nullable
NSString_inst_ObservationInfo(void* o) {
	void* _Nullable ret;
	@autoreleasepool {
		ret = [(NSString*)o observationInfo];
	}
	return ret;

}
void
NSString_inst_GetParagraphStart(void* o, void* startPtr, void* parEndPtr, void* contentsEndPtr, NSRange range) {
	@autoreleasepool {
		[(NSString*)o getParagraphStart:startPtr end:parEndPtr contentsEnd:contentsEndPtr forRange:range];
	}
}
void
NSString_inst_Dealloc(void* o) {
	@autoreleasepool {
		[(NSString*)o dealloc];
	}
}
void* _Nonnull
NSString_inst_ValuesForKeys(void* o, void* keys) {
	NSDictionary* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o valuesForKeys:keys];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void* _Nonnull
NSString_inst_StringByAppendingString(void* o, void* aString) {
	NSString* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o stringByAppendingString:aString];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void* _Nullable
NSString_inst_InitWithContentsOfFile(void* o, void* path) {
	NSObject* _Nullable ret;
	@autoreleasepool {
		ret = [(NSString*)o initWithContentsOfFile:path];
	}
	return ret;

}
void* _Nullable
NSString_inst_InitWithContentsOfFileEncoding(void* o, void* path, NSStringEncoding enc, void** error) {
	NSString* _Nullable ret;
	@autoreleasepool {
		ret = [(NSString*)o initWithContentsOfFile:path encoding:enc error:(NSError* _Nullable* _Nullable)error];
		for(int i=0;i<1;i++) {
			if(error[i] == 0) { break; }
			[(id)error[i] retain];
		}
	
	}
	return ret;

}
void* _Nullable
NSString_inst_InitWithContentsOfFileUsedEncoding(void* o, void* path, void* enc, void** error) {
	NSString* _Nullable ret;
	@autoreleasepool {
		ret = [(NSString*)o initWithContentsOfFile:path usedEncoding:enc error:(NSError* _Nullable* _Nullable)error];
		for(int i=0;i<1;i++) {
			if(error[i] == 0) { break; }
			[(id)error[i] retain];
		}
	
	}
	return ret;

}
NSComparisonResult
NSString_inst_CaseInsensitiveCompare(void* o, void* string) {
	NSComparisonResult ret;
	@autoreleasepool {
		ret = [(NSString*)o caseInsensitiveCompare:string];
	}
	return ret;

}
void* _Nonnull
NSString_inst_AttributeKeys(void* o) {
	NSArray* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o attributeKeys];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void* _Nonnull
NSString_inst_PathComponents(void* o) {
	NSArray* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o pathComponents];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
NSRange
NSString_inst_ParagraphRangeForRange(void* o, NSRange range) {
	NSRange ret;
	@autoreleasepool {
		ret = [(NSString*)o paragraphRangeForRange:range];
	}
	return ret;

}
void* _Nullable
NSString_inst_StringByAddingPercentEncodingWithAllowedCharacters(void* o, void* allowedCharacters) {
	NSString* _Nullable ret;
	@autoreleasepool {
		ret = [(NSString*)o stringByAddingPercentEncodingWithAllowedCharacters:allowedCharacters];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void* _Nullable
NSString_inst_InitWithContentsOfURL(void* o, void* url) {
	NSObject* _Nullable ret;
	@autoreleasepool {
		ret = [(NSString*)o initWithContentsOfURL:url];
	}
	return ret;

}
void* _Nullable
NSString_inst_InitWithContentsOfURLEncoding(void* o, void* url, NSStringEncoding enc, void** error) {
	NSString* _Nullable ret;
	@autoreleasepool {
		ret = [(NSString*)o initWithContentsOfURL:url encoding:enc error:(NSError* _Nullable* _Nullable)error];
		for(int i=0;i<1;i++) {
			if(error[i] == 0) { break; }
			[(id)error[i] retain];
		}
	
	}
	return ret;

}
void* _Nullable
NSString_inst_InitWithContentsOfURLUsedEncoding(void* o, void* url, void* enc, void** error) {
	NSString* _Nullable ret;
	@autoreleasepool {
		ret = [(NSString*)o initWithContentsOfURL:url usedEncoding:enc error:(NSError* _Nullable* _Nullable)error];
		for(int i=0;i<1;i++) {
			if(error[i] == 0) { break; }
			[(id)error[i] retain];
		}
	
	}
	return ret;

}
void
NSString_inst_GetLineStart(void* o, void* startPtr, void* lineEndPtr, void* contentsEndPtr, NSRange range) {
	@autoreleasepool {
		[(NSString*)o getLineStart:startPtr end:lineEndPtr contentsEnd:contentsEndPtr forRange:range];
	}
}
const void* _Nullable
NSString_inst_LossyCString(void* o) {
	const char* _Nullable ret;
	@autoreleasepool {
		ret = strdup([(NSString*)o lossyCString]);
	}
	return ret;

}
NSComparisonResult
NSString_inst_LocalizedCompare(void* o, void* string) {
	NSComparisonResult ret;
	@autoreleasepool {
		ret = [(NSString*)o localizedCompare:string];
	}
	return ret;

}
void
NSString_inst_HandleTakeValue(void* o, void* value, void* key) {
	@autoreleasepool {
		[(NSString*)o handleTakeValue:value forUnboundKey:key];
	}
}
const void* _Nonnull
NSString_inst_FileSystemRepresentation(void* o) {
	const char* _Nonnull ret;
	@autoreleasepool {
		ret = strdup([(NSString*)o fileSystemRepresentation]);
	}
	return ret;

}
void* _Nullable
NSString_inst_ClassForKeyedArchiver(void* o) {
	Class _Nullable ret;
	@autoreleasepool {
		ret = [(NSString*)o classForKeyedArchiver];
	}
	return ret;

}
void* _Nonnull
NSString_inst_LocalizedUppercaseString(void* o) {
	NSString* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o localizedUppercaseString];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void* _Nonnull
NSString_inst_CopyWithZone(void* o, void* zone) {
	NSObject* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o copyWithZone:zone];
	}
	return ret;

}
NSUInteger
NSString_inst_LengthOfBytesUsingEncoding(void* o, NSStringEncoding enc) {
	NSUInteger ret;
	@autoreleasepool {
		ret = [(NSString*)o lengthOfBytesUsingEncoding:enc];
	}
	return ret;

}
void* _Nullable
NSString_inst_InitWithData(void* o, void* data, NSStringEncoding encoding) {
	NSString* _Nullable ret;
	@autoreleasepool {
		ret = [(NSString*)o initWithData:data encoding:encoding];
	}
	return ret;

}
void* _Nullable
NSString_inst_PropertyListFromStringsFileFormat(void* o) {
	NSDictionary* _Nullable ret;
	@autoreleasepool {
		ret = [(NSString*)o propertyListFromStringsFileFormat];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void
NSString_inst_Release(void* o) {
	@autoreleasepool {
		[(NSString*)o release];
	}
}
BOOL
NSString_inst_IsKindOfClass(void* o, void* aClass) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSString*)o isKindOfClass:aClass];
	}
	return ret;

}
BOOL
NSString_inst_IsCaseInsensitiveLike(void* o, void* object) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSString*)o isCaseInsensitiveLike:object];
	}
	return ret;

}
void* _Nullable
NSString_inst_ValueWithUniqueID(void* o, void* uniqueID, void* key) {
	NSObject* _Nullable ret;
	@autoreleasepool {
		ret = [(NSString*)o valueWithUniqueID:uniqueID inPropertyWithKey:key];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void
NSString_inst_WillChange(void* o, NSKeyValueChange changeKind, void* indexes, void* key) {
	@autoreleasepool {
		[(NSString*)o willChange:changeKind valuesAtIndexes:indexes forKey:key];
	}
}
void* _Nullable
NSString_inst_InitWithCString(void* o, void* bytes) {
	NSObject* _Nullable ret;
	@autoreleasepool {
		ret = [(NSString*)o initWithCString:bytes];
	}
	return ret;

}
void* _Nullable
NSString_inst_InitWithCStringEncoding(void* o, void* nullTerminatedCString, NSStringEncoding encoding) {
	NSString* _Nullable ret;
	@autoreleasepool {
		ret = [(NSString*)o initWithCString:nullTerminatedCString encoding:encoding];
	}
	return ret;

}
void* _Nullable
NSString_inst_InitWithCStringLength(void* o, void* bytes, NSUInteger length) {
	NSObject* _Nullable ret;
	@autoreleasepool {
		ret = [(NSString*)o initWithCString:bytes length:length];
	}
	return ret;

}
void* _Nonnull
NSString_inst_WritableTypeIdentifiersForItemProvider(void* o) {
	NSArray* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o writableTypeIdentifiersForItemProvider];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
BOOL
NSString_inst_IsEqualToString(void* o, void* aString) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSString*)o isEqualToString:aString];
	}
	return ret;

}
void
NSString_inst_GetCharacters(void* o, void* buffer) {
	@autoreleasepool {
		[(NSString*)o getCharacters:buffer];
	}
}
void
NSString_inst_GetCharactersRange(void* o, void* buffer, NSRange range) {
	@autoreleasepool {
		[(NSString*)o getCharacters:buffer range:range];
	}
}
BOOL
NSString_inst_IsLike(void* o, void* object) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSString*)o isLike:object];
	}
	return ret;

}
void*
NSString_inst_Self(void* o) {
	NSString* ret;
	@autoreleasepool {
		ret = [(NSString*)o self];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
BOOL
NSString_inst_RespondsToSelector(void* o, void* aSelector) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSString*)o respondsToSelector:aSelector];
	}
	return ret;

}
BOOL
NSString_inst_WriteToFileAtomically(void* o, void* path, BOOL useAuxiliaryFile) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSString*)o writeToFile:path atomically:useAuxiliaryFile];
	}
	return ret;

}
BOOL
NSString_inst_WriteToFileAtomicallyEncoding(void* o, void* path, BOOL useAuxiliaryFile, NSStringEncoding enc, void** error) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSString*)o writeToFile:path atomically:useAuxiliaryFile encoding:enc error:(NSError* _Nullable* _Nullable)error];
		for(int i=0;i<1;i++) {
			if(error[i] == 0) { break; }
			[(id)error[i] retain];
		}
	
	}
	return ret;

}
void* _Nullable
NSString_inst_ValueForUndefinedKey(void* o, void* key) {
	NSObject* _Nullable ret;
	@autoreleasepool {
		ret = [(NSString*)o valueForUndefinedKey:key];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
BOOL
NSString_inst_IsLessThanOrEqualTo(void* o, void* object) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSString*)o isLessThanOrEqualTo:object];
	}
	return ret;

}
void* _Nonnull
NSString_inst_StringByFoldingWithOptions(void* o, NSStringCompareOptions options, void* locale) {
	NSString* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o stringByFoldingWithOptions:options locale:locale];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void* _Nonnull
NSString_inst_LocalizedCapitalizedString(void* o) {
	NSString* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o localizedCapitalizedString];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void* _Nullable
NSString_inst_StringByAppendingPathExtension(void* o, void* str) {
	NSString* _Nullable ret;
	@autoreleasepool {
		ret = [(NSString*)o stringByAppendingPathExtension:str];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void*
NSString_inst_Copy(void* o) {
	NSObject* ret;
	@autoreleasepool {
		ret = [(NSString*)o copy];
	}
	return ret;

}
void* _Nonnull
NSString_inst_StringsByAppendingPaths(void* o, void* paths) {
	NSArray* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o stringsByAppendingPaths:paths];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void* _Nullable
NSString_inst_ClassForArchiver(void* o) {
	Class _Nullable ret;
	@autoreleasepool {
		ret = [(NSString*)o classForArchiver];
	}
	return ret;

}
void
NSString_inst_URLResourceDataDidBecomeAvailable(void* o, void* sender, void* newBytes) {
	@autoreleasepool {
		[(NSString*)o URL:sender resourceDataDidBecomeAvailable:newBytes];
	}
}
void
NSString_inst_URLResourceDidFailLoadingWithReason(void* o, void* sender, void* reason) {
	@autoreleasepool {
		[(NSString*)o URL:sender resourceDidFailLoadingWithReason:reason];
	}
}
BOOL
NSString_inst_ScriptingIsEqualTo(void* o, void* object) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSString*)o scriptingIsEqualTo:object];
	}
	return ret;

}
void* _Nullable
NSString_inst_ObjectSpecifier(void* o) {
	NSScriptObjectSpecifier* _Nullable ret;
	@autoreleasepool {
		ret = [(NSString*)o objectSpecifier];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
BOOL
NSString_inst_IsNotEqualTo(void* o, void* object) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSString*)o isNotEqualTo:object];
	}
	return ret;

}
void* _Nonnull
NSString_inst_PrecomposedStringWithCompatibilityMapping(void* o) {
	NSString* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o precomposedStringWithCompatibilityMapping];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void* _Nonnull
NSString_inst_StringByPaddingToLength(void* o, NSUInteger newLength, void* padString, NSUInteger padIndex) {
	NSString* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o stringByPaddingToLength:newLength withString:padString startingAtIndex:padIndex];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void* _Nonnull
NSString_inst_SubstringFromIndex(void* o, NSUInteger from) {
	NSString* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o substringFromIndex:from];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void* _Nullable
NSString_inst_ScriptingProperties(void* o) {
	NSDictionary* _Nullable ret;
	@autoreleasepool {
		ret = [(NSString*)o scriptingProperties];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
BOOL
NSString_inst_IsGreaterThanOrEqualTo(void* o, void* object) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSString*)o isGreaterThanOrEqualTo:object];
	}
	return ret;

}
void* _Nullable
NSString_inst_InitWithBytes(void* o, void* bytes, NSUInteger len, NSStringEncoding encoding) {
	NSString* _Nullable ret;
	@autoreleasepool {
		ret = [(NSString*)o initWithBytes:bytes length:len encoding:encoding];
	}
	return ret;

}
void
NSString_inst_DidChange(void* o, NSKeyValueChange changeKind, void* indexes, void* key) {
	@autoreleasepool {
		[(NSString*)o didChange:changeKind valuesAtIndexes:indexes forKey:key];
	}
}
void
NSString_inst_ReplaceValueAtIndex(void* o, NSUInteger index, void* key, void* value) {
	@autoreleasepool {
		[(NSString*)o replaceValueAtIndex:index inPropertyWithKey:key withValue:value];
	}
}
void
NSString_inst_TakeStoredValue(void* o, void* value, void* key) {
	@autoreleasepool {
		[(NSString*)o takeStoredValue:value forKey:key];
	}
}
void
NSString_inst_SetNilValueForKey(void* o, void* key) {
	@autoreleasepool {
		[(NSString*)o setNilValueForKey:key];
	}
}
NSUInteger
NSString_inst_Hash(void* o) {
	NSUInteger ret;
	@autoreleasepool {
		ret = [(NSString*)o hash];
	}
	return ret;

}
NSComparisonResult
NSString_inst_LocalizedCaseInsensitiveCompare(void* o, void* string) {
	NSComparisonResult ret;
	@autoreleasepool {
		ret = [(NSString*)o localizedCaseInsensitiveCompare:string];
	}
	return ret;

}
void* _Nonnull
NSString_inst_MutableSetValueForKeyPath(void* o, void* keyPath) {
	NSMutableSet* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o mutableSetValueForKeyPath:keyPath];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
BOOL
NSString_inst_IsGreaterThan(void* o, void* object) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSString*)o isGreaterThan:object];
	}
	return ret;

}
BOOL
NSString_inst_IsAbsolutePath(void* o) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSString*)o isAbsolutePath];
	}
	return ret;

}
BOOL
NSString_inst_WriteToURLAtomically(void* o, void* url, BOOL atomically) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSString*)o writeToURL:url atomically:atomically];
	}
	return ret;

}
BOOL
NSString_inst_WriteToURLAtomicallyEncoding(void* o, void* url, BOOL useAuxiliaryFile, NSStringEncoding enc, void** error) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSString*)o writeToURL:url atomically:useAuxiliaryFile encoding:enc error:(NSError* _Nullable* _Nullable)error];
		for(int i=0;i<1;i++) {
			if(error[i] == 0) { break; }
			[(id)error[i] retain];
		}
	
	}
	return ret;

}
void
NSString_inst_SetObservationInfo(void* o, void* observationInfo) {
	@autoreleasepool {
		[(NSString*)o setObservationInfo:observationInfo];
	}
}
void*
NSString_inst_DebugDescription(void* o) {
	NSString* ret;
	@autoreleasepool {
		ret = [(NSString*)o debugDescription];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void* _Nullable
NSString_inst_StoredValueForKey(void* o, void* key) {
	NSObject* _Nullable ret;
	@autoreleasepool {
		ret = [(NSString*)o storedValueForKey:key];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void
NSString_inst_SetValueForKey(void* o, void* value, void* key) {
	@autoreleasepool {
		[(NSString*)o setValue:value forKey:key];
	}
}
void
NSString_inst_SetValueForKeyPath(void* o, void* value, void* keyPath) {
	@autoreleasepool {
		[(NSString*)o setValue:value forKeyPath:keyPath];
	}
}
void
NSString_inst_SetValueForUndefinedKey(void* o, void* value, void* key) {
	@autoreleasepool {
		[(NSString*)o setValue:value forUndefinedKey:key];
	}
}
BOOL
NSString_inst_DoesContain(void* o, void* object) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSString*)o doesContain:object];
	}
	return ret;

}
BOOL
NSString_inst_HasPrefix(void* o, void* str) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSString*)o hasPrefix:str];
	}
	return ret;

}
void* _Nonnull
NSString_inst_SubstringWithRange(void* o, NSRange range) {
	NSString* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o substringWithRange:range];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void* _Nullable
NSString_inst_InitWithCStringNoCopy(void* o, void* bytes, NSUInteger length, BOOL freeBuffer) {
	NSObject* _Nullable ret;
	@autoreleasepool {
		ret = [(NSString*)o initWithCStringNoCopy:bytes length:length freeWhenDone:freeBuffer];
	}
	return ret;

}
BOOL
NSString_inst_ValidateValueForKey(void* o, void** ioValue, void* inKey, void** outError) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSString*)o validateValue:(id _Nullable* _Nonnull)ioValue forKey:inKey error:(NSError* _Nullable* _Nullable)outError];
		for(int i=0;i<1;i++) {
			if(ioValue[i] == 0) { break; }
			[(id)ioValue[i] retain];
		}
	
	
		for(int i=0;i<1;i++) {
			if(outError[i] == 0) { break; }
			[(id)outError[i] retain];
		}
	
	}
	return ret;

}
BOOL
NSString_inst_ValidateValueForKeyPath(void* o, void** ioValue, void* inKeyPath, void** outError) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSString*)o validateValue:(id _Nullable* _Nonnull)ioValue forKeyPath:inKeyPath error:(NSError* _Nullable* _Nullable)outError];
		for(int i=0;i<1;i++) {
			if(ioValue[i] == 0) { break; }
			[(id)ioValue[i] retain];
		}
	
	
		for(int i=0;i<1;i++) {
			if(outError[i] == 0) { break; }
			[(id)outError[i] retain];
		}
	
	}
	return ret;

}
const void* _Nullable
NSString_inst_CString(void* o) {
	const char* _Nullable ret;
	@autoreleasepool {
		ret = strdup([(NSString*)o cString]);
	}
	return ret;

}
BOOL
NSString_inst_IsLessThan(void* o, void* object) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSString*)o isLessThan:object];
	}
	return ret;

}
void
NSString_inst_DidChangeValueForKey(void* o, void* key) {
	@autoreleasepool {
		[(NSString*)o didChangeValueForKey:key];
	}
}
void
NSString_inst_DidChangeValueForKeyWithSetMutation(void* o, void* key, NSKeyValueSetMutationKind mutationKind, void* objects) {
	@autoreleasepool {
		[(NSString*)o didChangeValueForKey:key withSetMutation:mutationKind usingObjects:objects];
	}
}
void* _Nullable
NSString_inst_CoerceValue(void* o, void* value, void* key) {
	NSObject* _Nullable ret;
	@autoreleasepool {
		ret = [(NSString*)o coerceValue:value forKey:key];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void* _Nonnull
NSString_inst_StringByAbbreviatingWithTildeInPath(void* o) {
	NSString* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o stringByAbbreviatingWithTildeInPath];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void* _Nonnull
NSString_inst_StringByStandardizingPath(void* o) {
	NSString* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o stringByStandardizingPath];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void
NSString_inst_SetScriptingProperties(void* o, void* scriptingProperties) {
	@autoreleasepool {
		[(NSString*)o setScriptingProperties:scriptingProperties];
	}
}
BOOL
NSString_inst_ScriptingBeginsWith(void* o, void* object) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSString*)o scriptingBeginsWith:object];
	}
	return ret;

}
NSRange
NSString_inst_RangeOfCharacterFromSet(void* o, void* searchSet) {
	NSRange ret;
	@autoreleasepool {
		ret = [(NSString*)o rangeOfCharacterFromSet:searchSet];
	}
	return ret;

}
NSRange
NSString_inst_RangeOfCharacterFromSetOptions(void* o, void* searchSet, NSStringCompareOptions mask) {
	NSRange ret;
	@autoreleasepool {
		ret = [(NSString*)o rangeOfCharacterFromSet:searchSet options:mask];
	}
	return ret;

}
NSRange
NSString_inst_RangeOfCharacterFromSetOptionsRange(void* o, void* searchSet, NSStringCompareOptions mask, NSRange rangeOfReceiverToSearch) {
	NSRange ret;
	@autoreleasepool {
		ret = [(NSString*)o rangeOfCharacterFromSet:searchSet options:mask range:rangeOfReceiverToSearch];
	}
	return ret;

}
void* _Nonnull
NSString_inst_CapitalizedString(void* o) {
	NSString* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o capitalizedString];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void* _Nonnull
NSString_inst_StringByReplacingCharactersInRange(void* o, NSRange range, void* replacement) {
	NSString* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o stringByReplacingCharactersInRange:range withString:replacement];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void
NSString_inst_RemoveObserverForKeyPath(void* o, void* observer, void* keyPath) {
	@autoreleasepool {
		[(NSString*)o removeObserver:observer forKeyPath:keyPath];
	}
}
void
NSString_inst_RemoveObserverForKeyPathContext(void* o, void* observer, void* keyPath, void* context) {
	@autoreleasepool {
		[(NSString*)o removeObserver:observer forKeyPath:keyPath context:context];
	}
}
void
NSString_inst_URLResourceDidFinishLoading(void* o, void* sender) {
	@autoreleasepool {
		[(NSString*)o URLResourceDidFinishLoading:sender];
	}
}
void* _Nullable
NSString_inst_AwakeAfterUsingCoder(void* o, void* aDecoder) {
	NSObject* _Nullable ret;
	@autoreleasepool {
		ret = [(NSString*)o awakeAfterUsingCoder:aDecoder];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void* _Nonnull
NSString_inst_MutableCopyWithZone(void* o, void* zone) {
	NSObject* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o mutableCopyWithZone:zone];
	}
	return ret;

}
void* _Nonnull
NSString_inst_PathExtension(void* o) {
	NSString* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o pathExtension];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
int
NSString_inst_IntValue(void* o) {
	int ret;
	@autoreleasepool {
		ret = [(NSString*)o intValue];
	}
	return ret;

}
NSStringEncoding
NSString_inst_SmallestEncoding(void* o) {
	NSStringEncoding ret;
	@autoreleasepool {
		ret = [(NSString*)o smallestEncoding];
	}
	return ret;

}
void
NSString_inst_ObserveValueForKeyPath(void* o, void* keyPath, void* object, void* change, void* context) {
	@autoreleasepool {
		[(NSString*)o observeValueForKeyPath:keyPath ofObject:object change:change context:context];
	}
}
void
NSString_inst_TakeValuesFromDictionary(void* o, void* properties) {
	@autoreleasepool {
		[(NSString*)o takeValuesFromDictionary:properties];
	}
}
void* _Nullable
NSString_inst_CopyScriptingValue(void* o, void* value, void* key, void* properties) {
	NSObject* _Nullable ret;
	@autoreleasepool {
		ret = [(NSString*)o copyScriptingValue:value forKey:key withProperties:properties];
	}
	return ret;

}
void* _Nonnull
NSString_inst_InitWithCharactersNoCopy(void* o, void* characters, NSUInteger length, BOOL freeBuffer) {
	NSString* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o initWithCharactersNoCopy:characters length:length freeWhenDone:freeBuffer];
	}
	return ret;

}
NSRange
NSString_inst_RangeOfComposedCharacterSequenceAtIndex(void* o, NSUInteger index) {
	NSRange ret;
	@autoreleasepool {
		ret = [(NSString*)o rangeOfComposedCharacterSequenceAtIndex:index];
	}
	return ret;

}
void* _Nullable
NSString_inst_StringByRemovingPercentEncoding(void* o) {
	NSString* _Nullable ret;
	@autoreleasepool {
		ret = [(NSString*)o stringByRemovingPercentEncoding];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void* _Nullable
NSString_inst_InverseForRelationshipKey(void* o, void* relationshipKey) {
	NSString* _Nullable ret;
	@autoreleasepool {
		ret = [(NSString*)o inverseForRelationshipKey:relationshipKey];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void* _Nonnull
NSString_inst_ToOneRelationshipKeys(void* o) {
	NSArray* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o toOneRelationshipKeys];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
BOOL
NSString_inst_IsProxy(void* o) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSString*)o isProxy];
	}
	return ret;

}
void* _Nonnull
NSString_inst_CommonPrefixWithString(void* o, void* str, NSStringCompareOptions mask) {
	NSString* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o commonPrefixWithString:str options:mask];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void* _Nullable
NSString_inst_StringByApplyingTransform(void* o, void* transform, BOOL reverse) {
	NSString* _Nullable ret;
	@autoreleasepool {
		ret = [(NSString*)o stringByApplyingTransform:transform reverse:reverse];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
BOOL
NSString_inst_ContainsString(void* o, void* str) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSString*)o containsString:str];
	}
	return ret;

}
void* _Nonnull
NSString_inst_LocalizedLowercaseString(void* o) {
	NSString* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o localizedLowercaseString];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void* _Nonnull
NSString_inst_Init(void* o) {
	NSString* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o init];
	}
	return ret;

}
void* _Nonnull
NSString_inst_MutableSetValueForKey(void* o, void* key) {
	NSMutableSet* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o mutableSetValueForKey:key];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
BOOL
NSString_inst_ScriptingContains(void* o, void* object) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSString*)o scriptingContains:object];
	}
	return ret;

}
void* _Nonnull
NSString_inst_InitWithCharacters(void* o, void* characters, NSUInteger length) {
	NSString* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o initWithCharacters:characters length:length];
	}
	return ret;

}
const void* _Nullable
NSString_inst_CStringUsingEncoding(void* o, NSStringEncoding encoding) {
	const char* _Nullable ret;
	@autoreleasepool {
		ret = strdup([(NSString*)o cStringUsingEncoding:encoding]);
	}
	return ret;

}
void
NSString_inst_InsertValueInPropertyWithKey(void* o, void* value, void* key) {
	@autoreleasepool {
		[(NSString*)o insertValue:value inPropertyWithKey:key];
	}
}
void
NSString_inst_InsertValueAtIndex(void* o, void* value, NSUInteger index, void* key) {
	@autoreleasepool {
		[(NSString*)o insertValue:value atIndex:index inPropertyWithKey:key];
	}
}
void* _Nonnull
NSString_inst_MutableArrayValueForKeyPath(void* o, void* keyPath) {
	NSMutableArray* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o mutableArrayValueForKeyPath:keyPath];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void*
NSString_inst_Class(void* o) {
	Class ret;
	@autoreleasepool {
		ret = [(NSString*)o class];
	}
	return ret;

}
BOOL
NSString_inst_CanBeConvertedToEncoding(void* o, NSStringEncoding encoding) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSString*)o canBeConvertedToEncoding:encoding];
	}
	return ret;

}
void* _Nonnull
NSString_inst_DecomposedStringWithCompatibilityMapping(void* o) {
	NSString* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o decomposedStringWithCompatibilityMapping];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
NSUInteger
NSString_inst_Length(void* o) {
	NSUInteger ret;
	@autoreleasepool {
		ret = [(NSString*)o length];
	}
	return ret;

}
void* _Nullable
NSString_inst_ReplacementObjectForKeyedArchiver(void* o, void* archiver) {
	NSObject* _Nullable ret;
	@autoreleasepool {
		ret = [(NSString*)o replacementObjectForKeyedArchiver:archiver];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void*
NSString_inst_Autorelease(void* o) {
	NSString* ret;
	@autoreleasepool {
		ret = [(NSString*)o autorelease];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
BOOL
NSString_inst_ScriptingIsLessThanOrEqualTo(void* o, void* object) {
	BOOL ret;
	@autoreleasepool {
		ret = [(NSString*)o scriptingIsLessThanOrEqualTo:object];
	}
	return ret;

}
void* _Nonnull
NSString_inst_InitWithString(void* o, void* aString) {
	NSString* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o initWithString:aString];
	}
	return ret;

}
void* _Nonnull
NSString_inst_LastPathComponent(void* o) {
	NSString* _Nonnull ret;
	@autoreleasepool {
		ret = [(NSString*)o lastPathComponent];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}
void* _Nullable
NSString_inst_ReplacementObjectForCoder(void* o, void* aCoder) {
	NSObject* _Nullable ret;
	@autoreleasepool {
		ret = [(NSString*)o replacementObjectForCoder:aCoder];
		if (ret != nil && ret != o) { [ret retain]; }
	}
	return ret;

}

void C1Dealloc(void*);
void C1Release(void*);

@interface c1 : NSObject 
{ }
- (void)dealloc;
- (void)release;
- (void)super_dealloc;
- (void)super_release;

@end
void c1_super_Dealloc(void* o);
void c1_super_Release(void* o);

@implementation c1

- (void)dealloc
{
	C1Dealloc(self);
	[super dealloc];
}


- (void)release
{
	C1Release(self);
}


- (void)super_dealloc
{
	[super dealloc];
}
		

- (void)super_release
{
	[super release];
}
		

@end

void c1_super_Dealloc(void* o)
{
	@autoreleasepool {
		[(c1*)o super_dealloc];
	}
}


void c1_super_Release(void* o)
{
	@autoreleasepool {
		[(c1*)o super_release];
	}
}


void*
c1Alloc() {
	return [c1 alloc];
}

void
c1_inst_Dealloc(void* o) {
	@autoreleasepool {
		[(c1*)o dealloc];
	}
}
void
c1_inst_Release(void* o) {
	@autoreleasepool {
		[(c1*)o release];
	}
}

void C2MyMethod(void*);

@interface c2 : NSObject 
{ }
- (void)myMethod;

@end

@implementation c2

- (void)myMethod
{
	C2MyMethod(self);
}



@end


void*
c2Alloc() {
	return [c2 alloc];
}

void
NSWrap_init() {
	[[NSThread new] start]; // put the runtime into multi-threaded mode
}

*/
import "C"

import (
	"unsafe"
	"runtime"
	"sync"
)

func init() {
	C.NSWrap_init()
}

type Id struct {
	ptr unsafe.Pointer
}
func (o *Id) Ptr() unsafe.Pointer { if o == nil { return nil }; return o.ptr }

type NSObject interface {
	Ptr() unsafe.Pointer
}

type _NSZone = C.struct__NSZone

type BOOL C.uchar

type Class *C.struct_objc_class

type NSMethodSignature struct { Id }
func (o *NSMethodSignature) Ptr() unsafe.Pointer { if o == nil { return nil }; return o.ptr }
func (o *Id) NSMethodSignature() *NSMethodSignature {
	return (*NSMethodSignature)(unsafe.Pointer(o))
}

type SEL *C.struct_objc_selector

type NSInteger C.long

type NSSet struct { Id }
func (o *NSSet) Ptr() unsafe.Pointer { if o == nil { return nil }; return o.ptr }
func (o *Id) NSSet() *NSSet {
	return (*NSSet)(unsafe.Pointer(o))
}

type NSString struct { Id }
func (o *NSString) Ptr() unsafe.Pointer { if o == nil { return nil }; return o.ptr }
func (o *Id) NSString() *NSString {
	return (*NSString)(unsafe.Pointer(o))
}

type Protocol interface {
	Ptr() unsafe.Pointer
}

type NSArray struct { Id }
func (o *NSArray) Ptr() unsafe.Pointer { if o == nil { return nil }; return o.ptr }
func (o *Id) NSArray() *NSArray {
	return (*NSArray)(unsafe.Pointer(o))
}

type NSUInteger C.ulong

type NSDictionary struct { Id }
func (o *NSDictionary) Ptr() unsafe.Pointer { if o == nil { return nil }; return o.ptr }
func (o *Id) NSDictionary() *NSDictionary {
	return (*NSDictionary)(unsafe.Pointer(o))
}

type NSCoder struct { Id }
func (o *NSCoder) Ptr() unsafe.Pointer { if o == nil { return nil }; return o.ptr }
func (o *Id) NSCoder() *NSCoder {
	return (*NSCoder)(unsafe.Pointer(o))
}

type NSPortCoder struct { NSCoder }
func (o *NSPortCoder) Ptr() unsafe.Pointer { if o == nil { return nil }; return o.ptr }
func (o *Id) NSPortCoder() *NSPortCoder {
	return (*NSPortCoder)(unsafe.Pointer(o))
}

type NSURL struct { Id }
func (o *NSURL) Ptr() unsafe.Pointer { if o == nil { return nil }; return o.ptr }
func (o *Id) NSURL() *NSURL {
	return (*NSURL)(unsafe.Pointer(o))
}

type NSData struct { Id }
func (o *NSData) Ptr() unsafe.Pointer { if o == nil { return nil }; return o.ptr }
func (o *Id) NSData() *NSData {
	return (*NSData)(unsafe.Pointer(o))
}

type NSKeyValueChange C.enum_NSKeyValueChange

type NSIndexSet struct { Id }
func (o *NSIndexSet) Ptr() unsafe.Pointer { if o == nil { return nil }; return o.ptr }
func (o *Id) NSIndexSet() *NSIndexSet {
	return (*NSIndexSet)(unsafe.Pointer(o))
}

type NSError struct { Id }
func (o *NSError) Ptr() unsafe.Pointer { if o == nil { return nil }; return o.ptr }
func (o *Id) NSError() *NSError {
	return (*NSError)(unsafe.Pointer(o))
}

type NSMutableSet struct { NSSet }
func (o *NSMutableSet) Ptr() unsafe.Pointer { if o == nil { return nil }; return o.ptr }
func (o *Id) NSMutableSet() *NSMutableSet {
	return (*NSMutableSet)(unsafe.Pointer(o))
}

type NSKeyValueSetMutationKind C.enum_NSKeyValueSetMutationKind

type NSInvocation struct { Id }
func (o *NSInvocation) Ptr() unsafe.Pointer { if o == nil { return nil }; return o.ptr }
func (o *Id) NSInvocation() *NSInvocation {
	return (*NSInvocation)(unsafe.Pointer(o))
}

type NSFileManager struct { Id }
func (o *NSFileManager) Ptr() unsafe.Pointer { if o == nil { return nil }; return o.ptr }
func (o *Id) NSFileManager() *NSFileManager {
	return (*NSFileManager)(unsafe.Pointer(o))
}

type NSOrderedSet struct { Id }
func (o *NSOrderedSet) Ptr() unsafe.Pointer { if o == nil { return nil }; return o.ptr }
func (o *Id) NSOrderedSet() *NSOrderedSet {
	return (*NSOrderedSet)(unsafe.Pointer(o))
}

type NSMutableOrderedSet struct { NSOrderedSet }
func (o *NSMutableOrderedSet) Ptr() unsafe.Pointer { if o == nil { return nil }; return o.ptr }
func (o *Id) NSMutableOrderedSet() *NSMutableOrderedSet {
	return (*NSMutableOrderedSet)(unsafe.Pointer(o))
}

type NSKeyValueObservingOptions C.enum_NSKeyValueObservingOptions

type NSClassDescription struct { Id }
func (o *NSClassDescription) Ptr() unsafe.Pointer { if o == nil { return nil }; return o.ptr }
func (o *Id) NSClassDescription() *NSClassDescription {
	return (*NSClassDescription)(unsafe.Pointer(o))
}

type FourCharCode C.UInt32

type NSScriptObjectSpecifier struct { Id }
func (o *NSScriptObjectSpecifier) Ptr() unsafe.Pointer { if o == nil { return nil }; return o.ptr }
func (o *Id) NSScriptObjectSpecifier() *NSScriptObjectSpecifier {
	return (*NSScriptObjectSpecifier)(unsafe.Pointer(o))
}

type NSKeyedArchiver struct { NSCoder }
func (o *NSKeyedArchiver) Ptr() unsafe.Pointer { if o == nil { return nil }; return o.ptr }
func (o *Id) NSKeyedArchiver() *NSKeyedArchiver {
	return (*NSKeyedArchiver)(unsafe.Pointer(o))
}

type NSMutableArray struct { NSArray }
func (o *NSMutableArray) Ptr() unsafe.Pointer { if o == nil { return nil }; return o.ptr }
func (o *Id) NSMutableArray() *NSMutableArray {
	return (*NSMutableArray)(unsafe.Pointer(o))
}

type NSTimeInterval C.double

type NSThread struct { Id }
func (o *NSThread) Ptr() unsafe.Pointer { if o == nil { return nil }; return o.ptr }
func (o *Id) NSThread() *NSThread {
	return (*NSThread)(unsafe.Pointer(o))
}

type NSArchiver struct { NSCoder }
func (o *NSArchiver) Ptr() unsafe.Pointer { if o == nil { return nil }; return o.ptr }
func (o *Id) NSArchiver() *NSArchiver {
	return (*NSArchiver)(unsafe.Pointer(o))
}

type Unichar C.ushort

type NSStringEncoding C.NSUInteger

type NSItemProviderRepresentationVisibility C.enum_NSItemProviderRepresentationVisibility

type Char C.char

type NSRange = C.struct__NSRange

type NSRangePointer *C.NSRange

type Double C.double

type Float C.float

type NSLinguisticTagScheme = *NSString

type NSLinguisticTaggerOptions C.enum_NSLinguisticTaggerOptions

type NSOrthography struct { Id }
func (o *NSOrthography) Ptr() unsafe.Pointer { if o == nil { return nil }; return o.ptr }
func (o *Id) NSOrthography() *NSOrthography {
	return (*NSOrthography)(unsafe.Pointer(o))
}

type NSCharacterSet struct { Id }
func (o *NSCharacterSet) Ptr() unsafe.Pointer { if o == nil { return nil }; return o.ptr }
func (o *Id) NSCharacterSet() *NSCharacterSet {
	return (*NSCharacterSet)(unsafe.Pointer(o))
}

type NSStringEncodingConversionOptions C.enum_NSStringEncodingConversionOptions

type NSLocale struct { Id }
func (o *NSLocale) Ptr() unsafe.Pointer { if o == nil { return nil }; return o.ptr }
func (o *Id) NSLocale() *NSLocale {
	return (*NSLocale)(unsafe.Pointer(o))
}

type NSStringCompareOptions C.enum_NSStringCompareOptions

type NSComparisonResult C.enum_NSComparisonResult

type LongLong C.longlong

type NSZone = C.struct__NSZone

type Int C.int

type NSStringTransform = *NSString

type C1 struct { Id }
func (o *C1) Ptr() unsafe.Pointer { if o == nil { return nil }; return o.ptr }
func (o *Id) C1() *C1 {
	return (*C1)(unsafe.Pointer(o))
}

type C2 struct { Id }
func (o *C2) Ptr() unsafe.Pointer { if o == nil { return nil }; return o.ptr }
func (o *Id) C2() *C2 {
	return (*C2)(unsafe.Pointer(o))
}

func Selector(s string) SEL {
	return (SEL)(unsafe.Pointer(C.selectorFromString(C.CString(s))))
}

func (o *NSString) String() string {
	utf8 := o.UTF8String()
	ret := utf8.String()
	utf8.Free()
	runtime.KeepAlive(o)
	return ret
}

func CharWithGoString(s string) *Char {
	return (*Char)(unsafe.Pointer(C.CString(s)))
}

func CharWithBytes(b []byte) *Char {
	return (*Char)(unsafe.Pointer(C.CString(string(b))))
}

func (c *Char) String() string {
	return C.GoString((*C.char)(c))
}

func (c *Char) Free() {
	C.free(unsafe.Pointer(c))
}

func NSObjectMutableCopyWithZone(zone *_NSZone) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSObject_MutableCopyWithZone(unsafe.Pointer(zone)))
	if ret.ptr == nil { return ret }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	return ret
}

func NSObjectNew() *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSObject_New())
	if ret.ptr == nil { return ret }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	return ret
}

func NSObjectIsSubclassOfClass(aClass Class) bool {
	ret := (C.NSObject_IsSubclassOfClass(unsafe.Pointer(aClass))) != 0
	return ret
}

func NSObjectInstanceMethodSignatureForSelector(aSelector SEL) *NSMethodSignature {
	ret := &NSMethodSignature{}
	ret.ptr = unsafe.Pointer(C.NSObject_InstanceMethodSignatureForSelector(unsafe.Pointer(aSelector)))
	if ret.ptr == nil { return ret }
	return ret
}

func NSObjectAlloc() *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSObject_Alloc())
	if ret.ptr == nil { return ret }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	return ret
}

func (o *Id) GC() {
	if o.ptr == nil { return }
	runtime.SetFinalizer(o, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
}

func NSObjectAccessInstanceVariablesDirectly() bool {
	ret := (C.NSObject_AccessInstanceVariablesDirectly()) != 0
	return ret
}

func NSObjectSetVersion(aVersion NSInteger)  {
	C.NSObject_SetVersion((C.NSInteger)(aVersion))
}

func NSObjectUseStoredAccessor() bool {
	ret := (C.NSObject_UseStoredAccessor()) != 0
	return ret
}

func NSObjectLoad()  {
	C.NSObject_Load()
}

func NSObjectKeyPathsForValuesAffectingValueForKey(key *NSString) *NSSet {
	ret := &NSSet{}
	ret.ptr = unsafe.Pointer(C.NSObject_KeyPathsForValuesAffectingValueForKey(key.Ptr()))
	if ret.ptr == nil { return ret }
	return ret
}

func NSObjectVersion() NSInteger {
	ret := (NSInteger)(C.NSObject_Version())
	return ret
}

func NSObjectConformsToProtocol(protocol Protocol) bool {
	ret := (C.NSObject_ConformsToProtocol(protocol.Ptr())) != 0
	return ret
}

func NSObjectCancelPreviousPerformRequestsWithTarget(aTarget NSObject)  {
	C.NSObject_CancelPreviousPerformRequestsWithTarget(aTarget.Ptr())
}

func NSObjectCancelPreviousPerformRequestsWithTargetSelector(aTarget NSObject, aSelector SEL, anArgument NSObject)  {
	C.NSObject_CancelPreviousPerformRequestsWithTargetSelector(aTarget.Ptr(), unsafe.Pointer(aSelector), anArgument.Ptr())
}

func NSObjectInstancesRespondToSelector(aSelector SEL) bool {
	ret := (C.NSObject_InstancesRespondToSelector(unsafe.Pointer(aSelector))) != 0
	return ret
}

func NSObjectDebugDescription() *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSObject_DebugDescription())
	if ret.ptr == nil { return ret }
	return ret
}

func NSObjectClass() Class {
	ret := (Class)(unsafe.Pointer(C.NSObject_Class()))
	return ret
}

func NSObjectClassFallbacksForKeyedArchiver() *NSArray {
	ret := &NSArray{}
	ret.ptr = unsafe.Pointer(C.NSObject_ClassFallbacksForKeyedArchiver())
	if ret.ptr == nil { return ret }
	return ret
}

func NSObjectSetKeys(keys *NSArray, dependentKey *NSString)  {
	C.NSObject_SetKeys(keys.Ptr(), dependentKey.Ptr())
}

func NSObjectResolveClassMethod(sel SEL) bool {
	ret := (C.NSObject_ResolveClassMethod(unsafe.Pointer(sel))) != 0
	return ret
}

func NSObjectSuperclass() Class {
	ret := (Class)(unsafe.Pointer(C.NSObject_Superclass()))
	return ret
}

func NSObjectAutomaticallyNotifiesObserversForKey(key *NSString) bool {
	ret := (C.NSObject_AutomaticallyNotifiesObserversForKey(key.Ptr())) != 0
	return ret
}

func NSObjectAllocWithZone(zone *_NSZone) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSObject_AllocWithZone(unsafe.Pointer(zone)))
	if ret.ptr == nil { return ret }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	return ret
}

func NSObjectResolveInstanceMethod(sel SEL) bool {
	ret := (C.NSObject_ResolveInstanceMethod(unsafe.Pointer(sel))) != 0
	return ret
}

func NSObjectClassForKeyedUnarchiver() Class {
	ret := (Class)(unsafe.Pointer(C.NSObject_ClassForKeyedUnarchiver()))
	return ret
}

func NSObjectDescription() *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSObject_Description())
	if ret.ptr == nil { return ret }
	return ret
}

func NSObjectCopyWithZone(zone *_NSZone) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSObject_CopyWithZone(unsafe.Pointer(zone)))
	if ret.ptr == nil { return ret }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	return ret
}

func NSObjectHash() NSUInteger {
	ret := (NSUInteger)(C.NSObject_Hash())
	return ret
}

func (o *Id) Description() *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSObject_inst_Description(o.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) MethodSignatureForSelector(aSelector SEL) *NSMethodSignature {
	ret := &NSMethodSignature{}
	ret.ptr = unsafe.Pointer(C.NSObject_inst_MethodSignatureForSelector(o.Ptr(), unsafe.Pointer(aSelector)))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSMethodSignature)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSMethodSignature) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) SetValuesForKeysWithDictionary(keyedValues *NSDictionary)  {
	C.NSObject_inst_SetValuesForKeysWithDictionary(o.Ptr(), keyedValues.Ptr())
	runtime.KeepAlive(o)
}

func (o *Id) ObserveValueForKeyPath(keyPath *NSString, object NSObject, change *NSDictionary, context unsafe.Pointer)  {
	C.NSObject_inst_ObserveValueForKeyPath(o.Ptr(), keyPath.Ptr(), object.Ptr(), change.Ptr(), unsafe.Pointer(context))
	runtime.KeepAlive(o)
}

func (o *Id) StoredValueForKey(key *NSString) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSObject_inst_StoredValueForKey(o.Ptr(), key.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) ClassForArchiver() Class {
	ret := (Class)(unsafe.Pointer(C.NSObject_inst_ClassForArchiver(o.Ptr())))
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) ReplacementObjectForPortCoder(coder *NSPortCoder) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSObject_inst_ReplacementObjectForPortCoder(o.Ptr(), coder.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) URLResourceDataDidBecomeAvailable(sender *NSURL, newBytes *NSData)  {
	C.NSObject_inst_URLResourceDataDidBecomeAvailable(o.Ptr(), sender.Ptr(), newBytes.Ptr())
	runtime.KeepAlive(o)
}

func (o *Id) URLResourceDidFailLoadingWithReason(sender *NSURL, reason *NSString)  {
	C.NSObject_inst_URLResourceDidFailLoadingWithReason(o.Ptr(), sender.Ptr(), reason.Ptr())
	runtime.KeepAlive(o)
}

func (o *Id) ValueForKeyPath(keyPath *NSString) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSObject_inst_ValueForKeyPath(o.Ptr(), keyPath.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) UnableToSetNilForKey(key *NSString)  {
	C.NSObject_inst_UnableToSetNilForKey(o.Ptr(), key.Ptr())
	runtime.KeepAlive(o)
}

func (o *Id) ScriptingIsGreaterThanOrEqualTo(object NSObject) bool {
	ret := (C.NSObject_inst_ScriptingIsGreaterThanOrEqualTo(o.Ptr(), object.Ptr())) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) ClassForPortCoder() Class {
	ret := (Class)(unsafe.Pointer(C.NSObject_inst_ClassForPortCoder(o.Ptr())))
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) InsertValueInPropertyWithKey(value NSObject, key *NSString)  {
	C.NSObject_inst_InsertValueInPropertyWithKey(o.Ptr(), value.Ptr(), key.Ptr())
	runtime.KeepAlive(o)
}

func (o *Id) InsertValueAtIndex(value NSObject, index NSUInteger, key *NSString)  {
	C.NSObject_inst_InsertValueAtIndex(o.Ptr(), value.Ptr(), (C.NSUInteger)(index), key.Ptr())
	runtime.KeepAlive(o)
}

func (o *Id) WillChange(changeKind NSKeyValueChange, indexes *NSIndexSet, key *NSString)  {
	C.NSObject_inst_WillChange(o.Ptr(), (C.NSKeyValueChange)(changeKind), indexes.Ptr(), key.Ptr())
	runtime.KeepAlive(o)
}

func (o *Id) AttemptRecoveryFromErrorOptionIndex(error *NSError, recoveryOptionIndex NSUInteger) bool {
	ret := (C.NSObject_inst_AttemptRecoveryFromErrorOptionIndex(o.Ptr(), error.Ptr(), (C.NSUInteger)(recoveryOptionIndex))) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) AttemptRecoveryFromErrorOptionIndexDelegate(error *NSError, recoveryOptionIndex NSUInteger, delegate NSObject, didRecoverSelector SEL, contextInfo unsafe.Pointer)  {
	C.NSObject_inst_AttemptRecoveryFromErrorOptionIndexDelegate(o.Ptr(), error.Ptr(), (C.NSUInteger)(recoveryOptionIndex), delegate.Ptr(), unsafe.Pointer(didRecoverSelector), unsafe.Pointer(contextInfo))
	runtime.KeepAlive(o)
}

func (o *Id) ScriptingIsLessThan(object NSObject) bool {
	ret := (C.NSObject_inst_ScriptingIsLessThan(o.Ptr(), object.Ptr())) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) Release()  {
	C.NSObject_inst_Release(o.Ptr())
	runtime.KeepAlive(o)
}

func (o *Id) NewScriptingObjectOfClass(objectClass Class, key *NSString, contentsValue NSObject, properties *NSDictionary) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSObject_inst_NewScriptingObjectOfClass(o.Ptr(), unsafe.Pointer(objectClass), key.Ptr(), contentsValue.Ptr(), properties.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) MutableSetValueForKeyPath(keyPath *NSString) *NSMutableSet {
	ret := &NSMutableSet{}
	ret.ptr = unsafe.Pointer(C.NSObject_inst_MutableSetValueForKeyPath(o.Ptr(), keyPath.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSMutableSet)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSMutableSet) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) WillChangeValueForKey(key *NSString)  {
	C.NSObject_inst_WillChangeValueForKey(o.Ptr(), key.Ptr())
	runtime.KeepAlive(o)
}

func (o *Id) WillChangeValueForKeyWithSetMutation(key *NSString, mutationKind NSKeyValueSetMutationKind, objects *NSSet)  {
	C.NSObject_inst_WillChangeValueForKeyWithSetMutation(o.Ptr(), key.Ptr(), (C.NSKeyValueSetMutationKind)(mutationKind), objects.Ptr())
	runtime.KeepAlive(o)
}

func (o *Id) Dealloc()  {
	C.NSObject_inst_Dealloc(o.Ptr())
	runtime.KeepAlive(o)
}

func (o *Id) Retain() *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSObject_inst_Retain(o.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) DidChange(changeKind NSKeyValueChange, indexes *NSIndexSet, key *NSString)  {
	C.NSObject_inst_DidChange(o.Ptr(), (C.NSKeyValueChange)(changeKind), indexes.Ptr(), key.Ptr())
	runtime.KeepAlive(o)
}

func (o *Id) IsLike(object *NSString) bool {
	ret := (C.NSObject_inst_IsLike(o.Ptr(), object.Ptr())) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) MutableSetValueForKey(key *NSString) *NSMutableSet {
	ret := &NSMutableSet{}
	ret.ptr = unsafe.Pointer(C.NSObject_inst_MutableSetValueForKey(o.Ptr(), key.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSMutableSet)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSMutableSet) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) IsKindOfClass(aClass Class) bool {
	ret := (C.NSObject_inst_IsKindOfClass(o.Ptr(), unsafe.Pointer(aClass))) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) ReplaceValueAtIndex(index NSUInteger, key *NSString, value NSObject)  {
	C.NSObject_inst_ReplaceValueAtIndex(o.Ptr(), (C.NSUInteger)(index), key.Ptr(), value.Ptr())
	runtime.KeepAlive(o)
}

func (o *Id) IsEqualTo(object NSObject) bool {
	ret := (C.NSObject_inst_IsEqualTo(o.Ptr(), object.Ptr())) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) DictionaryWithValuesForKeys(keys *NSArray) *NSDictionary {
	ret := &NSDictionary{}
	ret.ptr = unsafe.Pointer(C.NSObject_inst_DictionaryWithValuesForKeys(o.Ptr(), keys.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSDictionary)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSDictionary) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) RetainCount() NSUInteger {
	ret := (NSUInteger)(C.NSObject_inst_RetainCount(o.Ptr()))
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) ForwardInvocation(anInvocation *NSInvocation)  {
	C.NSObject_inst_ForwardInvocation(o.Ptr(), anInvocation.Ptr())
	runtime.KeepAlive(o)
}

func (o *Id) TakeValueForKey(value NSObject, key *NSString)  {
	C.NSObject_inst_TakeValueForKey(o.Ptr(), value.Ptr(), key.Ptr())
	runtime.KeepAlive(o)
}

func (o *Id) TakeValueForKeyPath(value NSObject, keyPath *NSString)  {
	C.NSObject_inst_TakeValueForKeyPath(o.Ptr(), value.Ptr(), keyPath.Ptr())
	runtime.KeepAlive(o)
}

func (o *Id) Copy() *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSObject_inst_Copy(o.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) URLResourceDidCancelLoading(sender *NSURL)  {
	C.NSObject_inst_URLResourceDidCancelLoading(o.Ptr(), sender.Ptr())
	runtime.KeepAlive(o)
}

func (o *Id) FileManagerShouldProceedAfterError(fm *NSFileManager, errorInfo *NSDictionary) bool {
	ret := (C.NSObject_inst_FileManagerShouldProceedAfterError(o.Ptr(), fm.Ptr(), errorInfo.Ptr())) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) FileManagerWillProcessPath(fm *NSFileManager, path *NSString)  {
	C.NSObject_inst_FileManagerWillProcessPath(o.Ptr(), fm.Ptr(), path.Ptr())
	runtime.KeepAlive(o)
}

func (o *Id) ConformsToProtocol(aProtocol Protocol) bool {
	ret := (C.NSObject_inst_ConformsToProtocol(o.Ptr(), aProtocol.Ptr())) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) ClassForKeyedArchiver() Class {
	ret := (Class)(unsafe.Pointer(C.NSObject_inst_ClassForKeyedArchiver(o.Ptr())))
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) SetValueForKey(value NSObject, key *NSString)  {
	C.NSObject_inst_SetValueForKey(o.Ptr(), value.Ptr(), key.Ptr())
	runtime.KeepAlive(o)
}

func (o *Id) SetValueForKeyPath(value NSObject, keyPath *NSString)  {
	C.NSObject_inst_SetValueForKeyPath(o.Ptr(), value.Ptr(), keyPath.Ptr())
	runtime.KeepAlive(o)
}

func (o *Id) SetValueForUndefinedKey(value NSObject, key *NSString)  {
	C.NSObject_inst_SetValueForUndefinedKey(o.Ptr(), value.Ptr(), key.Ptr())
	runtime.KeepAlive(o)
}

func (o *Id) IsMemberOfClass(aClass Class) bool {
	ret := (C.NSObject_inst_IsMemberOfClass(o.Ptr(), unsafe.Pointer(aClass))) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) SetObservationInfo(observationInfo unsafe.Pointer)  {
	C.NSObject_inst_SetObservationInfo(o.Ptr(), unsafe.Pointer(observationInfo))
	runtime.KeepAlive(o)
}

func (o *Id) ToManyRelationshipKeys() *NSArray {
	ret := &NSArray{}
	ret.ptr = unsafe.Pointer(C.NSObject_inst_ToManyRelationshipKeys(o.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSArray)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSArray) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) RemoveObserverForKeyPath(observer NSObject, keyPath *NSString)  {
	C.NSObject_inst_RemoveObserverForKeyPath(o.Ptr(), observer.Ptr(), keyPath.Ptr())
	runtime.KeepAlive(o)
}

func (o *Id) RemoveObserverForKeyPathContext(observer NSObject, keyPath *NSString, context unsafe.Pointer)  {
	C.NSObject_inst_RemoveObserverForKeyPathContext(o.Ptr(), observer.Ptr(), keyPath.Ptr(), unsafe.Pointer(context))
	runtime.KeepAlive(o)
}

func (o *Id) MutableOrderedSetValueForKeyPath(keyPath *NSString) *NSMutableOrderedSet {
	ret := &NSMutableOrderedSet{}
	ret.ptr = unsafe.Pointer(C.NSObject_inst_MutableOrderedSetValueForKeyPath(o.Ptr(), keyPath.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSMutableOrderedSet)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSMutableOrderedSet) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) AddObserver(observer NSObject, keyPath *NSString, options NSKeyValueObservingOptions, context unsafe.Pointer)  {
	C.NSObject_inst_AddObserver(o.Ptr(), observer.Ptr(), keyPath.Ptr(), (C.NSKeyValueObservingOptions)(options), unsafe.Pointer(context))
	runtime.KeepAlive(o)
}

func (o *Id) DoesNotRecognizeSelector(aSelector SEL)  {
	C.NSObject_inst_DoesNotRecognizeSelector(o.Ptr(), unsafe.Pointer(aSelector))
	runtime.KeepAlive(o)
}

func (o *Id) ValidateValueForKey(ioValue *[]*Id, inKey *NSString, outError *[]*NSError) bool {

	goSlice1 := make([]unsafe.Pointer,cap(*ioValue))
	for i := 0; i < len(*ioValue); i++ {
		goSlice1[i] = (*ioValue)[i].Ptr()
	}

	goSlice3 := make([]unsafe.Pointer,cap(*outError))
	for i := 0; i < len(*outError); i++ {
		goSlice3[i] = (*outError)[i].Ptr()
	}
	ret := (C.NSObject_inst_ValidateValueForKey(o.Ptr(), (*unsafe.Pointer)(unsafe.Pointer(&goSlice1[0])), inKey.Ptr(), (*unsafe.Pointer)(unsafe.Pointer(&goSlice3[0])))) != 0
	(*ioValue) = (*ioValue)[:cap(*ioValue)]
	for i := 0; i < len(*ioValue); i++ {
		if goSlice1[i] == nil {
			(*ioValue) = (*ioValue)[:i]
			break
		}
		if (*ioValue)[i] == nil {
			(*ioValue)[i] = &Id{}
			runtime.SetFinalizer((*ioValue)[i], func(o *Id) {
				o.Release()
			})
		}
		(*ioValue)[i].ptr = goSlice1[i]
	}
	(*outError) = (*outError)[:cap(*outError)]
	for i := 0; i < len(*outError); i++ {
		if goSlice3[i] == nil {
			(*outError) = (*outError)[:i]
			break
		}
		if (*outError)[i] == nil {
			(*outError)[i] = &NSError{}
			runtime.SetFinalizer((*outError)[i], func(o *NSError) {
				o.Release()
			})
		}
		(*outError)[i].ptr = goSlice3[i]
	}
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) ValidateValueForKeyPath(ioValue *[]*Id, inKeyPath *NSString, outError *[]*NSError) bool {

	goSlice1 := make([]unsafe.Pointer,cap(*ioValue))
	for i := 0; i < len(*ioValue); i++ {
		goSlice1[i] = (*ioValue)[i].Ptr()
	}

	goSlice3 := make([]unsafe.Pointer,cap(*outError))
	for i := 0; i < len(*outError); i++ {
		goSlice3[i] = (*outError)[i].Ptr()
	}
	ret := (C.NSObject_inst_ValidateValueForKeyPath(o.Ptr(), (*unsafe.Pointer)(unsafe.Pointer(&goSlice1[0])), inKeyPath.Ptr(), (*unsafe.Pointer)(unsafe.Pointer(&goSlice3[0])))) != 0
	(*ioValue) = (*ioValue)[:cap(*ioValue)]
	for i := 0; i < len(*ioValue); i++ {
		if goSlice1[i] == nil {
			(*ioValue) = (*ioValue)[:i]
			break
		}
		if (*ioValue)[i] == nil {
			(*ioValue)[i] = &Id{}
			runtime.SetFinalizer((*ioValue)[i], func(o *Id) {
				o.Release()
			})
		}
		(*ioValue)[i].ptr = goSlice1[i]
	}
	(*outError) = (*outError)[:cap(*outError)]
	for i := 0; i < len(*outError); i++ {
		if goSlice3[i] == nil {
			(*outError) = (*outError)[:i]
			break
		}
		if (*outError)[i] == nil {
			(*outError)[i] = &NSError{}
			runtime.SetFinalizer((*outError)[i], func(o *NSError) {
				o.Release()
			})
		}
		(*outError)[i].ptr = goSlice3[i]
	}
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) Self() *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSObject_inst_Self(o.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) ScriptingProperties() *NSDictionary {
	ret := &NSDictionary{}
	ret.ptr = unsafe.Pointer(C.NSObject_inst_ScriptingProperties(o.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSDictionary)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSDictionary) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) MutableCopy() *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSObject_inst_MutableCopy(o.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) ValueAtIndex(index NSUInteger, key *NSString) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSObject_inst_ValueAtIndex(o.Ptr(), (C.NSUInteger)(index), key.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) ClassDescription() *NSClassDescription {
	ret := &NSClassDescription{}
	ret.ptr = unsafe.Pointer(C.NSObject_inst_ClassDescription(o.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSClassDescription)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSClassDescription) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) TakeStoredValue(value NSObject, key *NSString)  {
	C.NSObject_inst_TakeStoredValue(o.Ptr(), value.Ptr(), key.Ptr())
	runtime.KeepAlive(o)
}

func (o *Id) IsCaseInsensitiveLike(object *NSString) bool {
	ret := (C.NSObject_inst_IsCaseInsensitiveLike(o.Ptr(), object.Ptr())) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) ValueWithName(name *NSString, key *NSString) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSObject_inst_ValueWithName(o.Ptr(), name.Ptr(), key.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) DoesContain(object NSObject) bool {
	ret := (C.NSObject_inst_DoesContain(o.Ptr(), object.Ptr())) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) ClassCode() FourCharCode {
	ret := (FourCharCode)(C.NSObject_inst_ClassCode(o.Ptr()))
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) ScriptingValueForSpecifier(objectSpecifier *NSScriptObjectSpecifier) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSObject_inst_ScriptingValueForSpecifier(o.Ptr(), objectSpecifier.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) IsGreaterThanOrEqualTo(object NSObject) bool {
	ret := (C.NSObject_inst_IsGreaterThanOrEqualTo(o.Ptr(), object.Ptr())) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) ScriptingIsEqualTo(object NSObject) bool {
	ret := (C.NSObject_inst_ScriptingIsEqualTo(o.Ptr(), object.Ptr())) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) ObjectSpecifier() *NSScriptObjectSpecifier {
	ret := &NSScriptObjectSpecifier{}
	ret.ptr = unsafe.Pointer(C.NSObject_inst_ObjectSpecifier(o.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSScriptObjectSpecifier)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSScriptObjectSpecifier) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) ValueForKey(key *NSString) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSObject_inst_ValueForKey(o.Ptr(), key.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) Superclass() Class {
	ret := (Class)(unsafe.Pointer(C.NSObject_inst_Superclass(o.Ptr())))
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) ReplacementObjectForKeyedArchiver(archiver *NSKeyedArchiver) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSObject_inst_ReplacementObjectForKeyedArchiver(o.Ptr(), archiver.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) HandleTakeValue(value NSObject, key *NSString)  {
	C.NSObject_inst_HandleTakeValue(o.Ptr(), value.Ptr(), key.Ptr())
	runtime.KeepAlive(o)
}

func (o *Id) URLResourceDidFinishLoading(sender *NSURL)  {
	C.NSObject_inst_URLResourceDidFinishLoading(o.Ptr(), sender.Ptr())
	runtime.KeepAlive(o)
}

func (o *Id) InverseForRelationshipKey(relationshipKey *NSString) *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSObject_inst_InverseForRelationshipKey(o.Ptr(), relationshipKey.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) SetNilValueForKey(key *NSString)  {
	C.NSObject_inst_SetNilValueForKey(o.Ptr(), key.Ptr())
	runtime.KeepAlive(o)
}

func (o *Id) Autorelease() *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSObject_inst_Autorelease(o.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) ToOneRelationshipKeys() *NSArray {
	ret := &NSArray{}
	ret.ptr = unsafe.Pointer(C.NSObject_inst_ToOneRelationshipKeys(o.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSArray)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSArray) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) IsLessThan(object NSObject) bool {
	ret := (C.NSObject_inst_IsLessThan(o.Ptr(), object.Ptr())) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) ForwardingTargetForSelector(aSelector SEL) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSObject_inst_ForwardingTargetForSelector(o.Ptr(), unsafe.Pointer(aSelector)))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) TakeValuesFromDictionary(properties *NSDictionary)  {
	C.NSObject_inst_TakeValuesFromDictionary(o.Ptr(), properties.Ptr())
	runtime.KeepAlive(o)
}

func (o *Id) SetScriptingProperties(scriptingProperties *NSDictionary)  {
	C.NSObject_inst_SetScriptingProperties(o.Ptr(), scriptingProperties.Ptr())
	runtime.KeepAlive(o)
}

func (o *Id) ObservationInfo() unsafe.Pointer {
	ret := (unsafe.Pointer)(unsafe.Pointer(C.NSObject_inst_ObservationInfo(o.Ptr())))
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) IsNotEqualTo(object NSObject) bool {
	ret := (C.NSObject_inst_IsNotEqualTo(o.Ptr(), object.Ptr())) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) HandleQueryWithUnboundKey(key *NSString) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSObject_inst_HandleQueryWithUnboundKey(o.Ptr(), key.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) CopyScriptingValue(value NSObject, key *NSString, properties *NSDictionary) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSObject_inst_CopyScriptingValue(o.Ptr(), value.Ptr(), key.Ptr(), properties.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) ValuesForKeys(keys *NSArray) *NSDictionary {
	ret := &NSDictionary{}
	ret.ptr = unsafe.Pointer(C.NSObject_inst_ValuesForKeys(o.Ptr(), keys.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSDictionary)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSDictionary) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) Init() *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSObject_inst_Init(o.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) ClassName() *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSObject_inst_ClassName(o.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) ScriptingIsLessThanOrEqualTo(object NSObject) bool {
	ret := (C.NSObject_inst_ScriptingIsLessThanOrEqualTo(o.Ptr(), object.Ptr())) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) IndicesOfObjectsByEvaluatingObjectSpecifier(specifier *NSScriptObjectSpecifier) *NSArray {
	ret := &NSArray{}
	ret.ptr = unsafe.Pointer(C.NSObject_inst_IndicesOfObjectsByEvaluatingObjectSpecifier(o.Ptr(), specifier.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSArray)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSArray) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) AutoContentAccessingProxy() *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSObject_inst_AutoContentAccessingProxy(o.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) ScriptingEndsWith(object NSObject) bool {
	ret := (C.NSObject_inst_ScriptingEndsWith(o.Ptr(), object.Ptr())) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) ScriptingIsGreaterThan(object NSObject) bool {
	ret := (C.NSObject_inst_ScriptingIsGreaterThan(o.Ptr(), object.Ptr())) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) MutableArrayValueForKey(key *NSString) *NSMutableArray {
	ret := &NSMutableArray{}
	ret.ptr = unsafe.Pointer(C.NSObject_inst_MutableArrayValueForKey(o.Ptr(), key.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSMutableArray)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSMutableArray) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) Zone() *_NSZone {
	ret := (*_NSZone)(unsafe.Pointer(C.NSObject_inst_Zone(o.Ptr())))
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) MutableArrayValueForKeyPath(keyPath *NSString) *NSMutableArray {
	ret := &NSMutableArray{}
	ret.ptr = unsafe.Pointer(C.NSObject_inst_MutableArrayValueForKeyPath(o.Ptr(), keyPath.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSMutableArray)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSMutableArray) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) ValueForUndefinedKey(key *NSString) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSObject_inst_ValueForUndefinedKey(o.Ptr(), key.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) IsProxy() bool {
	ret := (C.NSObject_inst_IsProxy(o.Ptr())) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) ScriptingContains(object NSObject) bool {
	ret := (C.NSObject_inst_ScriptingContains(o.Ptr(), object.Ptr())) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) PerformSelector(aSelector SEL) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSObject_inst_PerformSelector(o.Ptr(), unsafe.Pointer(aSelector)))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) PerformSelectorWithObject(aSelector SEL, object NSObject) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSObject_inst_PerformSelectorWithObject(o.Ptr(), unsafe.Pointer(aSelector), object.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) PerformSelectorWithObjectAfterDelay(aSelector SEL, anArgument NSObject, delay NSTimeInterval)  {
	C.NSObject_inst_PerformSelectorWithObjectAfterDelay(o.Ptr(), unsafe.Pointer(aSelector), anArgument.Ptr(), (C.NSTimeInterval)(delay))
	runtime.KeepAlive(o)
}

func (o *Id) PerformSelectorWithObjectWithObject(aSelector SEL, object1 NSObject, object2 NSObject) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSObject_inst_PerformSelectorWithObjectWithObject(o.Ptr(), unsafe.Pointer(aSelector), object1.Ptr(), object2.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) PerformSelectorWithObjectAfterDelayInModes(aSelector SEL, anArgument NSObject, delay NSTimeInterval, modes *NSArray)  {
	C.NSObject_inst_PerformSelectorWithObjectAfterDelayInModes(o.Ptr(), unsafe.Pointer(aSelector), anArgument.Ptr(), (C.NSTimeInterval)(delay), modes.Ptr())
	runtime.KeepAlive(o)
}

func (o *Id) PerformSelectorOnThread(aSelector SEL, thr *NSThread, arg NSObject, wait BOOL)  {
	C.NSObject_inst_PerformSelectorOnThread(o.Ptr(), unsafe.Pointer(aSelector), thr.Ptr(), arg.Ptr(), (C.BOOL)(wait))
	runtime.KeepAlive(o)
}

func (o *Id) PerformSelectorOnThreadWithObject(aSelector SEL, thr *NSThread, arg NSObject, wait BOOL, array *NSArray)  {
	C.NSObject_inst_PerformSelectorOnThreadWithObject(o.Ptr(), unsafe.Pointer(aSelector), thr.Ptr(), arg.Ptr(), (C.BOOL)(wait), array.Ptr())
	runtime.KeepAlive(o)
}

func (o *Id) RemoveValueAtIndex(index NSUInteger, key *NSString)  {
	C.NSObject_inst_RemoveValueAtIndex(o.Ptr(), (C.NSUInteger)(index), key.Ptr())
	runtime.KeepAlive(o)
}

func (o *Id) Hash() NSUInteger {
	ret := (NSUInteger)(C.NSObject_inst_Hash(o.Ptr()))
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) GetClass() Class {
	ret := (Class)(unsafe.Pointer(C.NSObject_inst_Class(o.Ptr())))
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) DebugDescription() *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSObject_inst_DebugDescription(o.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) AwakeAfterUsingCoder(aDecoder *NSCoder) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSObject_inst_AwakeAfterUsingCoder(o.Ptr(), aDecoder.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) DidChangeValueForKey(key *NSString)  {
	C.NSObject_inst_DidChangeValueForKey(o.Ptr(), key.Ptr())
	runtime.KeepAlive(o)
}

func (o *Id) DidChangeValueForKeyWithSetMutation(key *NSString, mutationKind NSKeyValueSetMutationKind, objects *NSSet)  {
	C.NSObject_inst_DidChangeValueForKeyWithSetMutation(o.Ptr(), key.Ptr(), (C.NSKeyValueSetMutationKind)(mutationKind), objects.Ptr())
	runtime.KeepAlive(o)
}

func (o *Id) MutableOrderedSetValueForKey(key *NSString) *NSMutableOrderedSet {
	ret := &NSMutableOrderedSet{}
	ret.ptr = unsafe.Pointer(C.NSObject_inst_MutableOrderedSetValueForKey(o.Ptr(), key.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSMutableOrderedSet)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSMutableOrderedSet) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) PerformSelectorInBackground(aSelector SEL, arg NSObject)  {
	C.NSObject_inst_PerformSelectorInBackground(o.Ptr(), unsafe.Pointer(aSelector), arg.Ptr())
	runtime.KeepAlive(o)
}

func (o *Id) ClassForCoder() Class {
	ret := (Class)(unsafe.Pointer(C.NSObject_inst_ClassForCoder(o.Ptr())))
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) ScriptingBeginsWith(object NSObject) bool {
	ret := (C.NSObject_inst_ScriptingBeginsWith(o.Ptr(), object.Ptr())) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) CoerceValue(value NSObject, key *NSString) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSObject_inst_CoerceValue(o.Ptr(), value.Ptr(), key.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) ReplacementObjectForArchiver(archiver *NSArchiver) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSObject_inst_ReplacementObjectForArchiver(o.Ptr(), archiver.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) IsEqual(object NSObject) bool {
	ret := (C.NSObject_inst_IsEqual(o.Ptr(), object.Ptr())) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) IsGreaterThan(object NSObject) bool {
	ret := (C.NSObject_inst_IsGreaterThan(o.Ptr(), object.Ptr())) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) AttributeKeys() *NSArray {
	ret := &NSArray{}
	ret.ptr = unsafe.Pointer(C.NSObject_inst_AttributeKeys(o.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSArray)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSArray) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) ReplacementObjectForCoder(aCoder *NSCoder) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSObject_inst_ReplacementObjectForCoder(o.Ptr(), aCoder.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) RespondsToSelector(aSelector SEL) bool {
	ret := (C.NSObject_inst_RespondsToSelector(o.Ptr(), unsafe.Pointer(aSelector))) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) ValueWithUniqueID(uniqueID NSObject, key *NSString) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSObject_inst_ValueWithUniqueID(o.Ptr(), uniqueID.Ptr(), key.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) IsLessThanOrEqualTo(object NSObject) bool {
	ret := (C.NSObject_inst_IsLessThanOrEqualTo(o.Ptr(), object.Ptr())) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *Id) PerformSelectorOnMainThreadWithObject(aSelector SEL, arg NSObject, wait BOOL)  {
	C.NSObject_inst_PerformSelectorOnMainThreadWithObject(o.Ptr(), unsafe.Pointer(aSelector), arg.Ptr(), (C.BOOL)(wait))
	runtime.KeepAlive(o)
}

func (o *Id) PerformSelectorOnMainThreadWithObjectWaitUntilDone(aSelector SEL, arg NSObject, wait BOOL, array *NSArray)  {
	C.NSObject_inst_PerformSelectorOnMainThreadWithObjectWaitUntilDone(o.Ptr(), unsafe.Pointer(aSelector), arg.Ptr(), (C.BOOL)(wait), array.Ptr())
	runtime.KeepAlive(o)
}

func NSStringWritableTypeIdentifiersForItemProvider() *NSArray {
	ret := &NSArray{}
	ret.ptr = unsafe.Pointer(C.NSString_WritableTypeIdentifiersForItemProvider())
	if ret.ptr == nil { return ret }
	return ret
}

func NSStringWithCharacters(characters *Unichar, length NSUInteger) *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_StringWithCharacters(unsafe.Pointer(characters), (C.NSUInteger)(length)))
	if ret.ptr == nil { return ret }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	return ret
}

func NSStringSetVersion(aVersion NSInteger)  {
	C.NSString_SetVersion((C.NSInteger)(aVersion))
}

func NSStringLoad()  {
	C.NSString_Load()
}

func NSStringSetKeys(keys *NSArray, dependentKey *NSString)  {
	C.NSString_SetKeys(keys.Ptr(), dependentKey.Ptr())
}

func NSStringSupportsSecureCoding() bool {
	ret := (C.NSString_SupportsSecureCoding()) != 0
	return ret
}

func NSStringWithFormat(format *NSString, objects ...NSObject) *NSString {
	var object [16]unsafe.Pointer
	for i,o := range objects {
		object[i] = o.Ptr()
	}
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_StringWithFormat(format.Ptr(), unsafe.Pointer(&object)))
	if ret.ptr == nil { return ret }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	return ret
}

func NSStringIsSubclassOfClass(aClass Class) bool {
	ret := (C.NSString_IsSubclassOfClass(unsafe.Pointer(aClass))) != 0
	return ret
}

func NSStringVersion() NSInteger {
	ret := (NSInteger)(C.NSString_Version())
	return ret
}

func NSStringAllocWithZone(zone *_NSZone) *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_AllocWithZone(unsafe.Pointer(zone)))
	if ret.ptr == nil { return ret }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	return ret
}

func NSStringSuperclass() Class {
	ret := (Class)(unsafe.Pointer(C.NSString_Superclass()))
	return ret
}

func NSStringAutomaticallyNotifiesObserversForKey(key *NSString) bool {
	ret := (C.NSString_AutomaticallyNotifiesObserversForKey(key.Ptr())) != 0
	return ret
}

func NSStringMutableCopyWithZone(zone *_NSZone) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSString_MutableCopyWithZone(unsafe.Pointer(zone)))
	if ret.ptr == nil { return ret }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	return ret
}

func NSStringNew() *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_New())
	if ret.ptr == nil { return ret }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	return ret
}

func NSStringCancelPreviousPerformRequestsWithTarget(aTarget NSObject)  {
	C.NSString_CancelPreviousPerformRequestsWithTarget(aTarget.Ptr())
}

func NSStringCancelPreviousPerformRequestsWithTargetSelector(aTarget NSObject, aSelector SEL, anArgument NSObject)  {
	C.NSString_CancelPreviousPerformRequestsWithTargetSelector(aTarget.Ptr(), unsafe.Pointer(aSelector), anArgument.Ptr())
}

func NSStringInstancesRespondToSelector(aSelector SEL) bool {
	ret := (C.NSString_InstancesRespondToSelector(unsafe.Pointer(aSelector))) != 0
	return ret
}

func NSStringAvailableStringEncodings() *NSStringEncoding {
	ret := (*NSStringEncoding)(unsafe.Pointer(C.NSString_AvailableStringEncodings()))
	return ret
}

func NSStringAccessInstanceVariablesDirectly() bool {
	ret := (C.NSString_AccessInstanceVariablesDirectly()) != 0
	return ret
}

func NSStringKeyPathsForValuesAffectingValueForKey(key *NSString) *NSSet {
	ret := &NSSet{}
	ret.ptr = unsafe.Pointer(C.NSString_KeyPathsForValuesAffectingValueForKey(key.Ptr()))
	if ret.ptr == nil { return ret }
	return ret
}

func NSStringString() *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_String())
	if ret.ptr == nil { return ret }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	return ret
}

func NSStringPathWithComponents(components *NSArray) *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_PathWithComponents(components.Ptr()))
	if ret.ptr == nil { return ret }
	return ret
}

func NSStringLocalizedStringWithFormat(format *NSString, objects ...NSObject) *NSString {
	var object [16]unsafe.Pointer
	for i,o := range objects {
		object[i] = o.Ptr()
	}
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_LocalizedStringWithFormat(format.Ptr(), unsafe.Pointer(&object)))
	if ret.ptr == nil { return ret }
	return ret
}

func NSStringWithContentsOfFile(path *NSString) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSString_StringWithContentsOfFile(path.Ptr()))
	if ret.ptr == nil { return ret }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	return ret
}

func NSStringWithContentsOfFileEncoding(path *NSString, enc NSStringEncoding, error *[]*NSError) *NSString {

	goSlice2 := make([]unsafe.Pointer,cap(*error))
	for i := 0; i < len(*error); i++ {
		goSlice2[i] = (*error)[i].Ptr()
	}
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_StringWithContentsOfFileEncoding(path.Ptr(), (C.NSStringEncoding)(enc), (*unsafe.Pointer)(unsafe.Pointer(&goSlice2[0]))))
	(*error) = (*error)[:cap(*error)]
	for i := 0; i < len(*error); i++ {
		if goSlice2[i] == nil {
			(*error) = (*error)[:i]
			break
		}
		if (*error)[i] == nil {
			(*error)[i] = &NSError{}
			runtime.SetFinalizer((*error)[i], func(o *NSError) {
				o.Release()
			})
		}
		(*error)[i].ptr = goSlice2[i]
	}
	if ret.ptr == nil { return ret }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	return ret
}

func NSStringWithContentsOfFileUsedEncoding(path *NSString, enc *NSStringEncoding, error *[]*NSError) *NSString {

	goSlice2 := make([]unsafe.Pointer,cap(*error))
	for i := 0; i < len(*error); i++ {
		goSlice2[i] = (*error)[i].Ptr()
	}
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_StringWithContentsOfFileUsedEncoding(path.Ptr(), unsafe.Pointer(enc), (*unsafe.Pointer)(unsafe.Pointer(&goSlice2[0]))))
	(*error) = (*error)[:cap(*error)]
	for i := 0; i < len(*error); i++ {
		if goSlice2[i] == nil {
			(*error) = (*error)[:i]
			break
		}
		if (*error)[i] == nil {
			(*error)[i] = &NSError{}
			runtime.SetFinalizer((*error)[i], func(o *NSError) {
				o.Release()
			})
		}
		(*error)[i].ptr = goSlice2[i]
	}
	if ret.ptr == nil { return ret }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	return ret
}

func NSStringItemProviderVisibilityForRepresentationWithTypeIdentifier(typeIdentifier *NSString) NSItemProviderRepresentationVisibility {
	ret := (NSItemProviderRepresentationVisibility)(C.NSString_ItemProviderVisibilityForRepresentationWithTypeIdentifier(typeIdentifier.Ptr()))
	return ret
}

func NSStringClass() Class {
	ret := (Class)(unsafe.Pointer(C.NSString_Class()))
	return ret
}

func NSStringHash() NSUInteger {
	ret := (NSUInteger)(C.NSString_Hash())
	return ret
}

func NSStringDefaultCStringEncoding() NSStringEncoding {
	ret := (NSStringEncoding)(C.NSString_DefaultCStringEncoding())
	return ret
}

func NSStringWithString(string *NSString) *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_StringWithString(string.Ptr()))
	if ret.ptr == nil { return ret }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	return ret
}

func NSStringWithGoString(string string) *NSString {
	string_chr := CharWithGoString(string)
	defer string_chr.Free()
	ret := NSStringWithString(NSStringWithUTF8String(string_chr))
	return ret
}

func NSStringUseStoredAccessor() bool {
	ret := (C.NSString_UseStoredAccessor()) != 0
	return ret
}

func NSStringConformsToProtocol(protocol Protocol) bool {
	ret := (C.NSString_ConformsToProtocol(protocol.Ptr())) != 0
	return ret
}

func NSStringEncodingForData(data *NSData, opts *NSDictionary, string *[]*NSString, usedLossyConversion *BOOL) NSStringEncoding {

	goSlice2 := make([]unsafe.Pointer,cap(*string))
	for i := 0; i < len(*string); i++ {
		goSlice2[i] = (*string)[i].Ptr()
	}
	ret := (NSStringEncoding)(C.NSString_StringEncodingForData(data.Ptr(), opts.Ptr(), (*unsafe.Pointer)(unsafe.Pointer(&goSlice2[0])), unsafe.Pointer(usedLossyConversion)))
	(*string) = (*string)[:cap(*string)]
	for i := 0; i < len(*string); i++ {
		if goSlice2[i] == nil {
			(*string) = (*string)[:i]
			break
		}
		if (*string)[i] == nil {
			(*string)[i] = &NSString{}
			runtime.SetFinalizer((*string)[i], func(o *NSString) {
				o.Release()
			})
		}
		(*string)[i].ptr = goSlice2[i]
	}
	return ret
}

func NSStringLocalizedNameOfStringEncoding(encoding NSStringEncoding) *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_LocalizedNameOfStringEncoding((C.NSStringEncoding)(encoding)))
	if ret.ptr == nil { return ret }
	return ret
}

func NSStringClassFallbacksForKeyedArchiver() *NSArray {
	ret := &NSArray{}
	ret.ptr = unsafe.Pointer(C.NSString_ClassFallbacksForKeyedArchiver())
	if ret.ptr == nil { return ret }
	return ret
}

func NSStringResolveInstanceMethod(sel SEL) bool {
	ret := (C.NSString_ResolveInstanceMethod(unsafe.Pointer(sel))) != 0
	return ret
}

func NSStringDescription() *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_Description())
	if ret.ptr == nil { return ret }
	return ret
}

func NSStringCopyWithZone(zone *_NSZone) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSString_CopyWithZone(unsafe.Pointer(zone)))
	if ret.ptr == nil { return ret }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	return ret
}

func NSStringReadableTypeIdentifiersForItemProvider() *NSArray {
	ret := &NSArray{}
	ret.ptr = unsafe.Pointer(C.NSString_ReadableTypeIdentifiersForItemProvider())
	if ret.ptr == nil { return ret }
	return ret
}

func NSStringWithUTF8String(nullTerminatedCString *Char) *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_StringWithUTF8String(unsafe.Pointer(nullTerminatedCString)))
	if ret.ptr == nil { return ret }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	return ret
}

func NSStringInstanceMethodSignatureForSelector(aSelector SEL) *NSMethodSignature {
	ret := &NSMethodSignature{}
	ret.ptr = unsafe.Pointer(C.NSString_InstanceMethodSignatureForSelector(unsafe.Pointer(aSelector)))
	if ret.ptr == nil { return ret }
	return ret
}

func NSStringClassForKeyedUnarchiver() Class {
	ret := (Class)(unsafe.Pointer(C.NSString_ClassForKeyedUnarchiver()))
	return ret
}

func NSStringResolveClassMethod(sel SEL) bool {
	ret := (C.NSString_ResolveClassMethod(unsafe.Pointer(sel))) != 0
	return ret
}

func NSStringObjectWithItemProviderData(data *NSData, typeIdentifier *NSString, outError *[]*NSError) *NSString {

	goSlice2 := make([]unsafe.Pointer,cap(*outError))
	for i := 0; i < len(*outError); i++ {
		goSlice2[i] = (*outError)[i].Ptr()
	}
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_ObjectWithItemProviderData(data.Ptr(), typeIdentifier.Ptr(), (*unsafe.Pointer)(unsafe.Pointer(&goSlice2[0]))))
	(*outError) = (*outError)[:cap(*outError)]
	for i := 0; i < len(*outError); i++ {
		if goSlice2[i] == nil {
			(*outError) = (*outError)[:i]
			break
		}
		if (*outError)[i] == nil {
			(*outError)[i] = &NSError{}
			runtime.SetFinalizer((*outError)[i], func(o *NSError) {
				o.Release()
			})
		}
		(*outError)[i].ptr = goSlice2[i]
	}
	if ret.ptr == nil { return ret }
	return ret
}

func NSStringWithCString(bytes *Char) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSString_StringWithCString(unsafe.Pointer(bytes)))
	if ret.ptr == nil { return ret }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	return ret
}

func NSStringWithCStringEncoding(cString *Char, enc NSStringEncoding) *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_StringWithCStringEncoding(unsafe.Pointer(cString), (C.NSStringEncoding)(enc)))
	if ret.ptr == nil { return ret }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	return ret
}

func NSStringWithCStringLength(bytes *Char, length NSUInteger) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSString_StringWithCStringLength(unsafe.Pointer(bytes), (C.NSUInteger)(length)))
	if ret.ptr == nil { return ret }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	return ret
}

func NSStringWithContentsOfURL(url *NSURL) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSString_StringWithContentsOfURL(url.Ptr()))
	if ret.ptr == nil { return ret }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	return ret
}

func NSStringWithContentsOfURLEncoding(url *NSURL, enc NSStringEncoding, error *[]*NSError) *NSString {

	goSlice2 := make([]unsafe.Pointer,cap(*error))
	for i := 0; i < len(*error); i++ {
		goSlice2[i] = (*error)[i].Ptr()
	}
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_StringWithContentsOfURLEncoding(url.Ptr(), (C.NSStringEncoding)(enc), (*unsafe.Pointer)(unsafe.Pointer(&goSlice2[0]))))
	(*error) = (*error)[:cap(*error)]
	for i := 0; i < len(*error); i++ {
		if goSlice2[i] == nil {
			(*error) = (*error)[:i]
			break
		}
		if (*error)[i] == nil {
			(*error)[i] = &NSError{}
			runtime.SetFinalizer((*error)[i], func(o *NSError) {
				o.Release()
			})
		}
		(*error)[i].ptr = goSlice2[i]
	}
	if ret.ptr == nil { return ret }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	return ret
}

func NSStringWithContentsOfURLUsedEncoding(url *NSURL, enc *NSStringEncoding, error *[]*NSError) *NSString {

	goSlice2 := make([]unsafe.Pointer,cap(*error))
	for i := 0; i < len(*error); i++ {
		goSlice2[i] = (*error)[i].Ptr()
	}
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_StringWithContentsOfURLUsedEncoding(url.Ptr(), unsafe.Pointer(enc), (*unsafe.Pointer)(unsafe.Pointer(&goSlice2[0]))))
	(*error) = (*error)[:cap(*error)]
	for i := 0; i < len(*error); i++ {
		if goSlice2[i] == nil {
			(*error) = (*error)[:i]
			break
		}
		if (*error)[i] == nil {
			(*error)[i] = &NSError{}
			runtime.SetFinalizer((*error)[i], func(o *NSError) {
				o.Release()
			})
		}
		(*error)[i].ptr = goSlice2[i]
	}
	if ret.ptr == nil { return ret }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	return ret
}

func NSStringAlloc() *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_Alloc())
	if ret.ptr == nil { return ret }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	return ret
}

func (o *NSString) GC() {
	if o.ptr == nil { return }
	runtime.SetFinalizer(o, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
}

func NSStringDebugDescription() *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_DebugDescription())
	if ret.ptr == nil { return ret }
	return ret
}

func (o *NSString) InitWithBytesNoCopy(bytes unsafe.Pointer, len_ NSUInteger, encoding NSStringEncoding, freeBuffer BOOL) *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_InitWithBytesNoCopy(o.Ptr(), unsafe.Pointer(bytes), (C.NSUInteger)(len_), (C.NSStringEncoding)(encoding), (C.BOOL)(freeBuffer)))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) FastestEncoding() NSStringEncoding {
	ret := (NSStringEncoding)(C.NSString_inst_FastestEncoding(o.Ptr()))
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) GetCString(bytes *Char)  {
	C.NSString_inst_GetCString(o.Ptr(), unsafe.Pointer(bytes))
	runtime.KeepAlive(o)
}

func (o *NSString) GetCStringMaxLength(bytes *Char, maxLength NSUInteger)  {
	C.NSString_inst_GetCStringMaxLength(o.Ptr(), unsafe.Pointer(bytes), (C.NSUInteger)(maxLength))
	runtime.KeepAlive(o)
}

func (o *NSString) GetCStringMaxLengthEncoding(buffer *Char, maxBufferCount NSUInteger, encoding NSStringEncoding) bool {
	ret := (C.NSString_inst_GetCStringMaxLengthEncoding(o.Ptr(), unsafe.Pointer(buffer), (C.NSUInteger)(maxBufferCount), (C.NSStringEncoding)(encoding))) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) GetCStringMaxLengthRange(bytes *Char, maxLength NSUInteger, aRange NSRange, leftoverRange NSRangePointer)  {
	C.NSString_inst_GetCStringMaxLengthRange(o.Ptr(), unsafe.Pointer(bytes), (C.NSUInteger)(maxLength), (C.NSRange)(aRange), unsafe.Pointer(leftoverRange))
	runtime.KeepAlive(o)
}

func (o *NSString) ValueWithName(name *NSString, key *NSString) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_ValueWithName(o.Ptr(), name.Ptr(), key.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) VariantFittingPresentationWidth(width NSInteger) *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_VariantFittingPresentationWidth(o.Ptr(), (C.NSInteger)(width)))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) ScriptingIsLessThan(object NSObject) bool {
	ret := (C.NSString_inst_ScriptingIsLessThan(o.Ptr(), object.Ptr())) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) ValueAtIndex(index NSUInteger, key *NSString) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_ValueAtIndex(o.Ptr(), (C.NSUInteger)(index), key.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) IndicesOfObjectsByEvaluatingObjectSpecifier(specifier *NSScriptObjectSpecifier) *NSArray {
	ret := &NSArray{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_IndicesOfObjectsByEvaluatingObjectSpecifier(o.Ptr(), specifier.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSArray)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSArray) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) MutableArrayValueForKey(key *NSString) *NSMutableArray {
	ret := &NSMutableArray{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_MutableArrayValueForKey(o.Ptr(), key.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSMutableArray)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSMutableArray) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) PerformSelector(aSelector SEL) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_PerformSelector(o.Ptr(), unsafe.Pointer(aSelector)))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) PerformSelectorWithObject(aSelector SEL, object NSObject) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_PerformSelectorWithObject(o.Ptr(), unsafe.Pointer(aSelector), object.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) PerformSelectorWithObjectAfterDelay(aSelector SEL, anArgument NSObject, delay NSTimeInterval)  {
	C.NSString_inst_PerformSelectorWithObjectAfterDelay(o.Ptr(), unsafe.Pointer(aSelector), anArgument.Ptr(), (C.NSTimeInterval)(delay))
	runtime.KeepAlive(o)
}

func (o *NSString) PerformSelectorWithObjectWithObject(aSelector SEL, object1 NSObject, object2 NSObject) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_PerformSelectorWithObjectWithObject(o.Ptr(), unsafe.Pointer(aSelector), object1.Ptr(), object2.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) PerformSelectorWithObjectAfterDelayInModes(aSelector SEL, anArgument NSObject, delay NSTimeInterval, modes *NSArray)  {
	C.NSString_inst_PerformSelectorWithObjectAfterDelayInModes(o.Ptr(), unsafe.Pointer(aSelector), anArgument.Ptr(), (C.NSTimeInterval)(delay), modes.Ptr())
	runtime.KeepAlive(o)
}

func (o *NSString) PerformSelectorOnThread(aSelector SEL, thr *NSThread, arg NSObject, wait BOOL)  {
	C.NSString_inst_PerformSelectorOnThread(o.Ptr(), unsafe.Pointer(aSelector), thr.Ptr(), arg.Ptr(), (C.BOOL)(wait))
	runtime.KeepAlive(o)
}

func (o *NSString) PerformSelectorOnThreadWithObject(aSelector SEL, thr *NSThread, arg NSObject, wait BOOL, array *NSArray)  {
	C.NSString_inst_PerformSelectorOnThreadWithObject(o.Ptr(), unsafe.Pointer(aSelector), thr.Ptr(), arg.Ptr(), (C.BOOL)(wait), array.Ptr())
	runtime.KeepAlive(o)
}

func (o *NSString) ClassForCoder() Class {
	ret := (Class)(unsafe.Pointer(C.NSString_inst_ClassForCoder(o.Ptr())))
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) DoubleValue() Double {
	ret := (Double)(C.NSString_inst_DoubleValue(o.Ptr()))
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) FloatValue() Float {
	ret := (Float)(C.NSString_inst_FloatValue(o.Ptr()))
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) ToManyRelationshipKeys() *NSArray {
	ret := &NSArray{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_ToManyRelationshipKeys(o.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSArray)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSArray) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) MutableOrderedSetValueForKeyPath(keyPath *NSString) *NSMutableOrderedSet {
	ret := &NSMutableOrderedSet{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_MutableOrderedSetValueForKeyPath(o.Ptr(), keyPath.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSMutableOrderedSet)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSMutableOrderedSet) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) ForwardingTargetForSelector(aSelector SEL) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_ForwardingTargetForSelector(o.Ptr(), unsafe.Pointer(aSelector)))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) Zone() *_NSZone {
	ret := (*_NSZone)(unsafe.Pointer(C.NSString_inst_Zone(o.Ptr())))
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) StringByDeletingPathExtension() *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_StringByDeletingPathExtension(o.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) LinguisticTagsInRange(range_ NSRange, scheme NSLinguisticTagScheme, options NSLinguisticTaggerOptions, orthography *NSOrthography, tokenRanges *[]*NSArray) *NSArray {

	goSlice5 := make([]unsafe.Pointer,cap(*tokenRanges))
	for i := 0; i < len(*tokenRanges); i++ {
		goSlice5[i] = (*tokenRanges)[i].Ptr()
	}
	ret := &NSArray{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_LinguisticTagsInRange(o.Ptr(), (C.NSRange)(range_), scheme.Ptr(), (C.NSLinguisticTaggerOptions)(options), orthography.Ptr(), (*unsafe.Pointer)(unsafe.Pointer(&goSlice5[0]))))
	(*tokenRanges) = (*tokenRanges)[:cap(*tokenRanges)]
	for i := 0; i < len(*tokenRanges); i++ {
		if goSlice5[i] == nil {
			(*tokenRanges) = (*tokenRanges)[:i]
			break
		}
		if (*tokenRanges)[i] == nil {
			(*tokenRanges)[i] = &NSArray{}
			runtime.SetFinalizer((*tokenRanges)[i], func(o *NSArray) {
				o.Release()
			})
		}
		(*tokenRanges)[i].ptr = goSlice5[i]
	}
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSArray)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSArray) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) InitWithCoder(aDecoder *NSCoder) *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_InitWithCoder(o.Ptr(), aDecoder.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) CStringLength() NSUInteger {
	ret := (NSUInteger)(C.NSString_inst_CStringLength(o.Ptr()))
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) StringByAppendingFormat(format *NSString, objects ...NSObject) *NSString {
	var object [16]unsafe.Pointer
	for i,o := range objects {
		object[i] = o.Ptr()
	}
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_StringByAppendingFormat(o.Ptr(), format.Ptr(), unsafe.Pointer(&object)))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) LocalizedStandardContainsString(str *NSString) bool {
	ret := (C.NSString_inst_LocalizedStandardContainsString(o.Ptr(), str.Ptr())) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) SubstringToIndex(to NSUInteger) *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_SubstringToIndex(o.Ptr(), (C.NSUInteger)(to)))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) SetValuesForKeysWithDictionary(keyedValues *NSDictionary)  {
	C.NSString_inst_SetValuesForKeysWithDictionary(o.Ptr(), keyedValues.Ptr())
	runtime.KeepAlive(o)
}

func (o *NSString) ValueForKeyPath(keyPath *NSString) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_ValueForKeyPath(o.Ptr(), keyPath.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) TakeValueForKey(value NSObject, key *NSString)  {
	C.NSString_inst_TakeValueForKey(o.Ptr(), value.Ptr(), key.Ptr())
	runtime.KeepAlive(o)
}

func (o *NSString) TakeValueForKeyPath(value NSObject, keyPath *NSString)  {
	C.NSString_inst_TakeValueForKeyPath(o.Ptr(), value.Ptr(), keyPath.Ptr())
	runtime.KeepAlive(o)
}

func (o *NSString) PrecomposedStringWithCanonicalMapping() *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_PrecomposedStringWithCanonicalMapping(o.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) ComponentsSeparatedByCharactersInSet(separator *NSCharacterSet) *NSArray {
	ret := &NSArray{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_ComponentsSeparatedByCharactersInSet(o.Ptr(), separator.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSArray)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSArray) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) DataUsingEncoding(encoding NSStringEncoding) *NSData {
	ret := &NSData{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_DataUsingEncoding(o.Ptr(), (C.NSStringEncoding)(encoding)))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSData)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSData) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) DataUsingEncodingAllowLossyConversion(encoding NSStringEncoding, lossy BOOL) *NSData {
	ret := &NSData{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_DataUsingEncodingAllowLossyConversion(o.Ptr(), (C.NSStringEncoding)(encoding), (C.BOOL)(lossy)))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSData)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSData) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) IsEqual(object NSObject) bool {
	ret := (C.NSString_inst_IsEqual(o.Ptr(), object.Ptr())) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) GetBytes(buffer unsafe.Pointer, maxBufferCount NSUInteger, usedBufferCount *NSUInteger, encoding NSStringEncoding, options NSStringEncodingConversionOptions, range_ NSRange, leftover NSRangePointer) bool {
	ret := (C.NSString_inst_GetBytes(o.Ptr(), unsafe.Pointer(buffer), (C.NSUInteger)(maxBufferCount), unsafe.Pointer(usedBufferCount), (C.NSStringEncoding)(encoding), (C.NSStringEncodingConversionOptions)(options), (C.NSRange)(range_), unsafe.Pointer(leftover))) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) CompletePathIntoString(outputName *[]*NSString, flag BOOL, outputArray *[]*NSArray, filterTypes *NSArray) NSUInteger {

	goSlice1 := make([]unsafe.Pointer,cap(*outputName))
	for i := 0; i < len(*outputName); i++ {
		goSlice1[i] = (*outputName)[i].Ptr()
	}

	goSlice3 := make([]unsafe.Pointer,cap(*outputArray))
	for i := 0; i < len(*outputArray); i++ {
		goSlice3[i] = (*outputArray)[i].Ptr()
	}
	ret := (NSUInteger)(C.NSString_inst_CompletePathIntoString(o.Ptr(), (*unsafe.Pointer)(unsafe.Pointer(&goSlice1[0])), (C.BOOL)(flag), (*unsafe.Pointer)(unsafe.Pointer(&goSlice3[0])), filterTypes.Ptr()))
	(*outputName) = (*outputName)[:cap(*outputName)]
	for i := 0; i < len(*outputName); i++ {
		if goSlice1[i] == nil {
			(*outputName) = (*outputName)[:i]
			break
		}
		if (*outputName)[i] == nil {
			(*outputName)[i] = &NSString{}
			runtime.SetFinalizer((*outputName)[i], func(o *NSString) {
				o.Release()
			})
		}
		(*outputName)[i].ptr = goSlice1[i]
	}
	(*outputArray) = (*outputArray)[:cap(*outputArray)]
	for i := 0; i < len(*outputArray); i++ {
		if goSlice3[i] == nil {
			(*outputArray) = (*outputArray)[:i]
			break
		}
		if (*outputArray)[i] == nil {
			(*outputArray)[i] = &NSArray{}
			runtime.SetFinalizer((*outputArray)[i], func(o *NSArray) {
				o.Release()
			})
		}
		(*outputArray)[i].ptr = goSlice3[i]
	}
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) ClassForPortCoder() Class {
	ret := (Class)(unsafe.Pointer(C.NSString_inst_ClassForPortCoder(o.Ptr())))
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) IsMemberOfClass(aClass Class) bool {
	ret := (C.NSString_inst_IsMemberOfClass(o.Ptr(), unsafe.Pointer(aClass))) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) StringByAppendingPathComponent(str *NSString) *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_StringByAppendingPathComponent(o.Ptr(), str.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) ScriptingEndsWith(object NSObject) bool {
	ret := (C.NSString_inst_ScriptingEndsWith(o.Ptr(), object.Ptr())) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) ReplacementObjectForArchiver(archiver *NSArchiver) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_ReplacementObjectForArchiver(o.Ptr(), archiver.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) BoolValue() bool {
	ret := (C.NSString_inst_BoolValue(o.Ptr())) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) ItemProviderVisibilityForRepresentationWithTypeIdentifier(typeIdentifier *NSString) NSItemProviderRepresentationVisibility {
	ret := (NSItemProviderRepresentationVisibility)(C.NSString_inst_ItemProviderVisibilityForRepresentationWithTypeIdentifier(o.Ptr(), typeIdentifier.Ptr()))
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) LocalizedCaseInsensitiveContainsString(str *NSString) bool {
	ret := (C.NSString_inst_LocalizedCaseInsensitiveContainsString(o.Ptr(), str.Ptr())) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) HandleQueryWithUnboundKey(key *NSString) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_HandleQueryWithUnboundKey(o.Ptr(), key.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) ScriptingIsGreaterThan(object NSObject) bool {
	ret := (C.NSString_inst_ScriptingIsGreaterThan(o.Ptr(), object.Ptr())) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) RemoveValueAtIndex(index NSUInteger, key *NSString)  {
	C.NSString_inst_RemoveValueAtIndex(o.Ptr(), (C.NSUInteger)(index), key.Ptr())
	runtime.KeepAlive(o)
}

func (o *NSString) MaximumLengthOfBytesUsingEncoding(enc NSStringEncoding) NSUInteger {
	ret := (NSUInteger)(C.NSString_inst_MaximumLengthOfBytesUsingEncoding(o.Ptr(), (C.NSStringEncoding)(enc)))
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) CapitalizedStringWithLocale(locale *NSLocale) *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_CapitalizedStringWithLocale(o.Ptr(), locale.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) UppercaseStringWithLocale(locale *NSLocale) *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_UppercaseStringWithLocale(o.Ptr(), locale.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) Description() *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_Description(o.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) IntegerValue() NSInteger {
	ret := (NSInteger)(C.NSString_inst_IntegerValue(o.Ptr()))
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) ReplacementObjectForPortCoder(coder *NSPortCoder) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_ReplacementObjectForPortCoder(o.Ptr(), coder.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) DoesNotRecognizeSelector(aSelector SEL)  {
	C.NSString_inst_DoesNotRecognizeSelector(o.Ptr(), unsafe.Pointer(aSelector))
	runtime.KeepAlive(o)
}

func (o *NSString) AutoContentAccessingProxy() *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_AutoContentAccessingProxy(o.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) InitWithUTF8String(nullTerminatedCString *Char) *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_InitWithUTF8String(o.Ptr(), unsafe.Pointer(nullTerminatedCString)))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) StringByAddingPercentEscapesUsingEncoding(enc NSStringEncoding) *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_StringByAddingPercentEscapesUsingEncoding(o.Ptr(), (C.NSStringEncoding)(enc)))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) ClassName() *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_ClassName(o.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) ComponentsSeparatedByString(separator *NSString) *NSArray {
	ret := &NSArray{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_ComponentsSeparatedByString(o.Ptr(), separator.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSArray)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSArray) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) ConformsToProtocol(aProtocol Protocol) bool {
	ret := (C.NSString_inst_ConformsToProtocol(o.Ptr(), aProtocol.Ptr())) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) ValueForKey(key *NSString) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_ValueForKey(o.Ptr(), key.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) StringByDeletingLastPathComponent() *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_StringByDeletingLastPathComponent(o.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) StringByTrimmingCharactersInSet(set *NSCharacterSet) *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_StringByTrimmingCharactersInSet(o.Ptr(), set.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) UnableToSetNilForKey(key *NSString)  {
	C.NSString_inst_UnableToSetNilForKey(o.Ptr(), key.Ptr())
	runtime.KeepAlive(o)
}

func (o *NSString) DictionaryWithValuesForKeys(keys *NSArray) *NSDictionary {
	ret := &NSDictionary{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_DictionaryWithValuesForKeys(o.Ptr(), keys.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSDictionary)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSDictionary) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) MutableCopy() *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_MutableCopy(o.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) RangeOfString(searchString *NSString) NSRange {
	ret := (NSRange)(C.NSString_inst_RangeOfString(o.Ptr(), searchString.Ptr()))
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) RangeOfStringOptions(searchString *NSString, mask NSStringCompareOptions) NSRange {
	ret := (NSRange)(C.NSString_inst_RangeOfStringOptions(o.Ptr(), searchString.Ptr(), (C.NSStringCompareOptions)(mask)))
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) RangeOfStringOptionsRange(searchString *NSString, mask NSStringCompareOptions, rangeOfReceiverToSearch NSRange) NSRange {
	ret := (NSRange)(C.NSString_inst_RangeOfStringOptionsRange(o.Ptr(), searchString.Ptr(), (C.NSStringCompareOptions)(mask), (C.NSRange)(rangeOfReceiverToSearch)))
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) RangeOfStringOptionsRangeLocale(searchString *NSString, mask NSStringCompareOptions, rangeOfReceiverToSearch NSRange, locale *NSLocale) NSRange {
	ret := (NSRange)(C.NSString_inst_RangeOfStringOptionsRangeLocale(o.Ptr(), searchString.Ptr(), (C.NSStringCompareOptions)(mask), (C.NSRange)(rangeOfReceiverToSearch), locale.Ptr()))
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) LowercaseString() *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_LowercaseString(o.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) Retain() *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_Retain(o.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) ForwardInvocation(anInvocation *NSInvocation)  {
	C.NSString_inst_ForwardInvocation(o.Ptr(), anInvocation.Ptr())
	runtime.KeepAlive(o)
}

func (o *NSString) ScriptingValueForSpecifier(objectSpecifier *NSScriptObjectSpecifier) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_ScriptingValueForSpecifier(o.Ptr(), objectSpecifier.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) PropertyList() *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_PropertyList(o.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) LocalizedStandardRangeOfString(str *NSString) NSRange {
	ret := (NSRange)(C.NSString_inst_LocalizedStandardRangeOfString(o.Ptr(), str.Ptr()))
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) MethodSignatureForSelector(aSelector SEL) *NSMethodSignature {
	ret := &NSMethodSignature{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_MethodSignatureForSelector(o.Ptr(), unsafe.Pointer(aSelector)))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSMethodSignature)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSMethodSignature) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) AddObserver(observer NSObject, keyPath *NSString, options NSKeyValueObservingOptions, context unsafe.Pointer)  {
	C.NSString_inst_AddObserver(o.Ptr(), observer.Ptr(), keyPath.Ptr(), (C.NSKeyValueObservingOptions)(options), unsafe.Pointer(context))
	runtime.KeepAlive(o)
}

func (o *NSString) GetFileSystemRepresentation(cname *Char, max NSUInteger) bool {
	ret := (C.NSString_inst_GetFileSystemRepresentation(o.Ptr(), unsafe.Pointer(cname), (C.NSUInteger)(max))) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) StringByReplacingOccurrencesOfStringWithString(target *NSString, replacement *NSString) *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_StringByReplacingOccurrencesOfStringWithString(o.Ptr(), target.Ptr(), replacement.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) StringByReplacingOccurrencesOfStringWithStringOptions(target *NSString, replacement *NSString, options NSStringCompareOptions, searchRange NSRange) *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_StringByReplacingOccurrencesOfStringWithStringOptions(o.Ptr(), target.Ptr(), replacement.Ptr(), (C.NSStringCompareOptions)(options), (C.NSRange)(searchRange)))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) UTF8String() *Char {
	ret := (*Char)(unsafe.Pointer(C.NSString_inst_UTF8String(o.Ptr())))
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) FileManagerShouldProceedAfterError(fm *NSFileManager, errorInfo *NSDictionary) bool {
	ret := (C.NSString_inst_FileManagerShouldProceedAfterError(o.Ptr(), fm.Ptr(), errorInfo.Ptr())) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) FileManagerWillProcessPath(fm *NSFileManager, path *NSString)  {
	C.NSString_inst_FileManagerWillProcessPath(o.Ptr(), fm.Ptr(), path.Ptr())
	runtime.KeepAlive(o)
}

func (o *NSString) LowercaseStringWithLocale(locale *NSLocale) *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_LowercaseStringWithLocale(o.Ptr(), locale.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) RangeOfComposedCharacterSequencesForRange(range_ NSRange) NSRange {
	ret := (NSRange)(C.NSString_inst_RangeOfComposedCharacterSequencesForRange(o.Ptr(), (C.NSRange)(range_)))
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) ScriptingIsGreaterThanOrEqualTo(object NSObject) bool {
	ret := (C.NSString_inst_ScriptingIsGreaterThanOrEqualTo(o.Ptr(), object.Ptr())) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) MutableOrderedSetValueForKey(key *NSString) *NSMutableOrderedSet {
	ret := &NSMutableOrderedSet{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_MutableOrderedSetValueForKey(o.Ptr(), key.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSMutableOrderedSet)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSMutableOrderedSet) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) PerformSelectorOnMainThreadWithObject(aSelector SEL, arg NSObject, wait BOOL)  {
	C.NSString_inst_PerformSelectorOnMainThreadWithObject(o.Ptr(), unsafe.Pointer(aSelector), arg.Ptr(), (C.BOOL)(wait))
	runtime.KeepAlive(o)
}

func (o *NSString) PerformSelectorOnMainThreadWithObjectWaitUntilDone(aSelector SEL, arg NSObject, wait BOOL, array *NSArray)  {
	C.NSString_inst_PerformSelectorOnMainThreadWithObjectWaitUntilDone(o.Ptr(), unsafe.Pointer(aSelector), arg.Ptr(), (C.BOOL)(wait), array.Ptr())
	runtime.KeepAlive(o)
}

func (o *NSString) DecomposedStringWithCanonicalMapping() *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_DecomposedStringWithCanonicalMapping(o.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) LineRangeForRange(range_ NSRange) NSRange {
	ret := (NSRange)(C.NSString_inst_LineRangeForRange(o.Ptr(), (C.NSRange)(range_)))
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) CharacterAtIndex(index NSUInteger) Unichar {
	ret := (Unichar)(C.NSString_inst_CharacterAtIndex(o.Ptr(), (C.NSUInteger)(index)))
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) ClassDescription() *NSClassDescription {
	ret := &NSClassDescription{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_ClassDescription(o.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSClassDescription)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSClassDescription) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) NewScriptingObjectOfClass(objectClass Class, key *NSString, contentsValue NSObject, properties *NSDictionary) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_NewScriptingObjectOfClass(o.Ptr(), unsafe.Pointer(objectClass), key.Ptr(), contentsValue.Ptr(), properties.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) URLResourceDidCancelLoading(sender *NSURL)  {
	C.NSString_inst_URLResourceDidCancelLoading(o.Ptr(), sender.Ptr())
	runtime.KeepAlive(o)
}

func (o *NSString) LocalizedStandardCompare(string *NSString) NSComparisonResult {
	ret := (NSComparisonResult)(C.NSString_inst_LocalizedStandardCompare(o.Ptr(), string.Ptr()))
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) StringByReplacingPercentEscapesUsingEncoding(enc NSStringEncoding) *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_StringByReplacingPercentEscapesUsingEncoding(o.Ptr(), (C.NSStringEncoding)(enc)))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) WillChangeValueForKey(key *NSString)  {
	C.NSString_inst_WillChangeValueForKey(o.Ptr(), key.Ptr())
	runtime.KeepAlive(o)
}

func (o *NSString) WillChangeValueForKeyWithSetMutation(key *NSString, mutationKind NSKeyValueSetMutationKind, objects *NSSet)  {
	C.NSString_inst_WillChangeValueForKeyWithSetMutation(o.Ptr(), key.Ptr(), (C.NSKeyValueSetMutationKind)(mutationKind), objects.Ptr())
	runtime.KeepAlive(o)
}

func (o *NSString) AttemptRecoveryFromErrorOptionIndex(error *NSError, recoveryOptionIndex NSUInteger) bool {
	ret := (C.NSString_inst_AttemptRecoveryFromErrorOptionIndex(o.Ptr(), error.Ptr(), (C.NSUInteger)(recoveryOptionIndex))) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) AttemptRecoveryFromErrorOptionIndexDelegate(error *NSError, recoveryOptionIndex NSUInteger, delegate NSObject, didRecoverSelector SEL, contextInfo unsafe.Pointer)  {
	C.NSString_inst_AttemptRecoveryFromErrorOptionIndexDelegate(o.Ptr(), error.Ptr(), (C.NSUInteger)(recoveryOptionIndex), delegate.Ptr(), unsafe.Pointer(didRecoverSelector), unsafe.Pointer(contextInfo))
	runtime.KeepAlive(o)
}

func (o *NSString) ClassCode() FourCharCode {
	ret := (FourCharCode)(C.NSString_inst_ClassCode(o.Ptr()))
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) Compare(string *NSString) NSComparisonResult {
	ret := (NSComparisonResult)(C.NSString_inst_Compare(o.Ptr(), string.Ptr()))
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) CompareOptions(string *NSString, mask NSStringCompareOptions) NSComparisonResult {
	ret := (NSComparisonResult)(C.NSString_inst_CompareOptions(o.Ptr(), string.Ptr(), (C.NSStringCompareOptions)(mask)))
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) CompareOptionsRange(string *NSString, mask NSStringCompareOptions, rangeOfReceiverToCompare NSRange) NSComparisonResult {
	ret := (NSComparisonResult)(C.NSString_inst_CompareOptionsRange(o.Ptr(), string.Ptr(), (C.NSStringCompareOptions)(mask), (C.NSRange)(rangeOfReceiverToCompare)))
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) CompareOptionsRangeLocale(string *NSString, mask NSStringCompareOptions, rangeOfReceiverToCompare NSRange, locale NSObject) NSComparisonResult {
	ret := (NSComparisonResult)(C.NSString_inst_CompareOptionsRangeLocale(o.Ptr(), string.Ptr(), (C.NSStringCompareOptions)(mask), (C.NSRange)(rangeOfReceiverToCompare), locale.Ptr()))
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) HasSuffix(str *NSString) bool {
	ret := (C.NSString_inst_HasSuffix(o.Ptr(), str.Ptr())) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) StringByResolvingSymlinksInPath() *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_StringByResolvingSymlinksInPath(o.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) InitWithFormat(format *NSString, objects ...NSObject) *NSString {
	var object [16]unsafe.Pointer
	for i,o := range objects {
		object[i] = o.Ptr()
	}
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_InitWithFormat(o.Ptr(), format.Ptr(), unsafe.Pointer(&object)))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) InitWithFormatLocale(format *NSString, locale NSObject, objects ...NSObject) *NSString {
	var object [16]unsafe.Pointer
	for i,o := range objects {
		object[i] = o.Ptr()
	}
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_InitWithFormatLocale(o.Ptr(), format.Ptr(), locale.Ptr(), unsafe.Pointer(&object)))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) LongLongValue() LongLong {
	ret := (LongLong)(C.NSString_inst_LongLongValue(o.Ptr()))
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) StringByExpandingTildeInPath() *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_StringByExpandingTildeInPath(o.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) Superclass() Class {
	ret := (Class)(unsafe.Pointer(C.NSString_inst_Superclass(o.Ptr())))
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) PerformSelectorInBackground(aSelector SEL, arg NSObject)  {
	C.NSString_inst_PerformSelectorInBackground(o.Ptr(), unsafe.Pointer(aSelector), arg.Ptr())
	runtime.KeepAlive(o)
}

func (o *NSString) UppercaseString() *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_UppercaseString(o.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) IsEqualTo(object NSObject) bool {
	ret := (C.NSString_inst_IsEqualTo(o.Ptr(), object.Ptr())) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) RetainCount() NSUInteger {
	ret := (NSUInteger)(C.NSString_inst_RetainCount(o.Ptr()))
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) ObservationInfo() unsafe.Pointer {
	ret := (unsafe.Pointer)(unsafe.Pointer(C.NSString_inst_ObservationInfo(o.Ptr())))
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) GetParagraphStart(startPtr *NSUInteger, parEndPtr *NSUInteger, contentsEndPtr *NSUInteger, range_ NSRange)  {
	C.NSString_inst_GetParagraphStart(o.Ptr(), unsafe.Pointer(startPtr), unsafe.Pointer(parEndPtr), unsafe.Pointer(contentsEndPtr), (C.NSRange)(range_))
	runtime.KeepAlive(o)
}

func (o *NSString) Dealloc()  {
	C.NSString_inst_Dealloc(o.Ptr())
	runtime.KeepAlive(o)
}

func (o *NSString) ValuesForKeys(keys *NSArray) *NSDictionary {
	ret := &NSDictionary{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_ValuesForKeys(o.Ptr(), keys.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSDictionary)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSDictionary) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) StringByAppendingString(aString *NSString) *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_StringByAppendingString(o.Ptr(), aString.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) InitWithContentsOfFile(path *NSString) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_InitWithContentsOfFile(o.Ptr(), path.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) InitWithContentsOfFileEncoding(path *NSString, enc NSStringEncoding, error *[]*NSError) *NSString {

	goSlice3 := make([]unsafe.Pointer,cap(*error))
	for i := 0; i < len(*error); i++ {
		goSlice3[i] = (*error)[i].Ptr()
	}
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_InitWithContentsOfFileEncoding(o.Ptr(), path.Ptr(), (C.NSStringEncoding)(enc), (*unsafe.Pointer)(unsafe.Pointer(&goSlice3[0]))))
	(*error) = (*error)[:cap(*error)]
	for i := 0; i < len(*error); i++ {
		if goSlice3[i] == nil {
			(*error) = (*error)[:i]
			break
		}
		if (*error)[i] == nil {
			(*error)[i] = &NSError{}
			runtime.SetFinalizer((*error)[i], func(o *NSError) {
				o.Release()
			})
		}
		(*error)[i].ptr = goSlice3[i]
	}
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) InitWithContentsOfFileUsedEncoding(path *NSString, enc *NSStringEncoding, error *[]*NSError) *NSString {

	goSlice3 := make([]unsafe.Pointer,cap(*error))
	for i := 0; i < len(*error); i++ {
		goSlice3[i] = (*error)[i].Ptr()
	}
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_InitWithContentsOfFileUsedEncoding(o.Ptr(), path.Ptr(), unsafe.Pointer(enc), (*unsafe.Pointer)(unsafe.Pointer(&goSlice3[0]))))
	(*error) = (*error)[:cap(*error)]
	for i := 0; i < len(*error); i++ {
		if goSlice3[i] == nil {
			(*error) = (*error)[:i]
			break
		}
		if (*error)[i] == nil {
			(*error)[i] = &NSError{}
			runtime.SetFinalizer((*error)[i], func(o *NSError) {
				o.Release()
			})
		}
		(*error)[i].ptr = goSlice3[i]
	}
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) CaseInsensitiveCompare(string *NSString) NSComparisonResult {
	ret := (NSComparisonResult)(C.NSString_inst_CaseInsensitiveCompare(o.Ptr(), string.Ptr()))
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) AttributeKeys() *NSArray {
	ret := &NSArray{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_AttributeKeys(o.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSArray)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSArray) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) PathComponents() *NSArray {
	ret := &NSArray{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_PathComponents(o.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSArray)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSArray) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) ParagraphRangeForRange(range_ NSRange) NSRange {
	ret := (NSRange)(C.NSString_inst_ParagraphRangeForRange(o.Ptr(), (C.NSRange)(range_)))
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) StringByAddingPercentEncodingWithAllowedCharacters(allowedCharacters *NSCharacterSet) *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_StringByAddingPercentEncodingWithAllowedCharacters(o.Ptr(), allowedCharacters.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) InitWithContentsOfURL(url *NSURL) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_InitWithContentsOfURL(o.Ptr(), url.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) InitWithContentsOfURLEncoding(url *NSURL, enc NSStringEncoding, error *[]*NSError) *NSString {

	goSlice3 := make([]unsafe.Pointer,cap(*error))
	for i := 0; i < len(*error); i++ {
		goSlice3[i] = (*error)[i].Ptr()
	}
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_InitWithContentsOfURLEncoding(o.Ptr(), url.Ptr(), (C.NSStringEncoding)(enc), (*unsafe.Pointer)(unsafe.Pointer(&goSlice3[0]))))
	(*error) = (*error)[:cap(*error)]
	for i := 0; i < len(*error); i++ {
		if goSlice3[i] == nil {
			(*error) = (*error)[:i]
			break
		}
		if (*error)[i] == nil {
			(*error)[i] = &NSError{}
			runtime.SetFinalizer((*error)[i], func(o *NSError) {
				o.Release()
			})
		}
		(*error)[i].ptr = goSlice3[i]
	}
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) InitWithContentsOfURLUsedEncoding(url *NSURL, enc *NSStringEncoding, error *[]*NSError) *NSString {

	goSlice3 := make([]unsafe.Pointer,cap(*error))
	for i := 0; i < len(*error); i++ {
		goSlice3[i] = (*error)[i].Ptr()
	}
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_InitWithContentsOfURLUsedEncoding(o.Ptr(), url.Ptr(), unsafe.Pointer(enc), (*unsafe.Pointer)(unsafe.Pointer(&goSlice3[0]))))
	(*error) = (*error)[:cap(*error)]
	for i := 0; i < len(*error); i++ {
		if goSlice3[i] == nil {
			(*error) = (*error)[:i]
			break
		}
		if (*error)[i] == nil {
			(*error)[i] = &NSError{}
			runtime.SetFinalizer((*error)[i], func(o *NSError) {
				o.Release()
			})
		}
		(*error)[i].ptr = goSlice3[i]
	}
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) GetLineStart(startPtr *NSUInteger, lineEndPtr *NSUInteger, contentsEndPtr *NSUInteger, range_ NSRange)  {
	C.NSString_inst_GetLineStart(o.Ptr(), unsafe.Pointer(startPtr), unsafe.Pointer(lineEndPtr), unsafe.Pointer(contentsEndPtr), (C.NSRange)(range_))
	runtime.KeepAlive(o)
}

func (o *NSString) LossyCString() *Char {
	ret := (*Char)(unsafe.Pointer(C.NSString_inst_LossyCString(o.Ptr())))
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) LocalizedCompare(string *NSString) NSComparisonResult {
	ret := (NSComparisonResult)(C.NSString_inst_LocalizedCompare(o.Ptr(), string.Ptr()))
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) HandleTakeValue(value NSObject, key *NSString)  {
	C.NSString_inst_HandleTakeValue(o.Ptr(), value.Ptr(), key.Ptr())
	runtime.KeepAlive(o)
}

func (o *NSString) FileSystemRepresentation() *Char {
	ret := (*Char)(unsafe.Pointer(C.NSString_inst_FileSystemRepresentation(o.Ptr())))
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) ClassForKeyedArchiver() Class {
	ret := (Class)(unsafe.Pointer(C.NSString_inst_ClassForKeyedArchiver(o.Ptr())))
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) LocalizedUppercaseString() *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_LocalizedUppercaseString(o.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) CopyWithZone(zone *NSZone) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_CopyWithZone(o.Ptr(), unsafe.Pointer(zone)))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) LengthOfBytesUsingEncoding(enc NSStringEncoding) NSUInteger {
	ret := (NSUInteger)(C.NSString_inst_LengthOfBytesUsingEncoding(o.Ptr(), (C.NSStringEncoding)(enc)))
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) InitWithData(data *NSData, encoding NSStringEncoding) *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_InitWithData(o.Ptr(), data.Ptr(), (C.NSStringEncoding)(encoding)))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) PropertyListFromStringsFileFormat() *NSDictionary {
	ret := &NSDictionary{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_PropertyListFromStringsFileFormat(o.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSDictionary)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSDictionary) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) Release()  {
	C.NSString_inst_Release(o.Ptr())
	runtime.KeepAlive(o)
}

func (o *NSString) IsKindOfClass(aClass Class) bool {
	ret := (C.NSString_inst_IsKindOfClass(o.Ptr(), unsafe.Pointer(aClass))) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) IsCaseInsensitiveLike(object *NSString) bool {
	ret := (C.NSString_inst_IsCaseInsensitiveLike(o.Ptr(), object.Ptr())) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) ValueWithUniqueID(uniqueID NSObject, key *NSString) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_ValueWithUniqueID(o.Ptr(), uniqueID.Ptr(), key.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) WillChange(changeKind NSKeyValueChange, indexes *NSIndexSet, key *NSString)  {
	C.NSString_inst_WillChange(o.Ptr(), (C.NSKeyValueChange)(changeKind), indexes.Ptr(), key.Ptr())
	runtime.KeepAlive(o)
}

func (o *NSString) InitWithCString(bytes *Char) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_InitWithCString(o.Ptr(), unsafe.Pointer(bytes)))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) InitWithCStringEncoding(nullTerminatedCString *Char, encoding NSStringEncoding) *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_InitWithCStringEncoding(o.Ptr(), unsafe.Pointer(nullTerminatedCString), (C.NSStringEncoding)(encoding)))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) InitWithCStringLength(bytes *Char, length NSUInteger) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_InitWithCStringLength(o.Ptr(), unsafe.Pointer(bytes), (C.NSUInteger)(length)))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) WritableTypeIdentifiersForItemProvider() *NSArray {
	ret := &NSArray{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_WritableTypeIdentifiersForItemProvider(o.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSArray)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSArray) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) IsEqualToString(aString *NSString) bool {
	ret := (C.NSString_inst_IsEqualToString(o.Ptr(), aString.Ptr())) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) GetCharacters(buffer *Unichar)  {
	C.NSString_inst_GetCharacters(o.Ptr(), unsafe.Pointer(buffer))
	runtime.KeepAlive(o)
}

func (o *NSString) GetCharactersRange(buffer *Unichar, range_ NSRange)  {
	C.NSString_inst_GetCharactersRange(o.Ptr(), unsafe.Pointer(buffer), (C.NSRange)(range_))
	runtime.KeepAlive(o)
}

func (o *NSString) IsLike(object *NSString) bool {
	ret := (C.NSString_inst_IsLike(o.Ptr(), object.Ptr())) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) Self() *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_Self(o.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) RespondsToSelector(aSelector SEL) bool {
	ret := (C.NSString_inst_RespondsToSelector(o.Ptr(), unsafe.Pointer(aSelector))) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) WriteToFileAtomically(path *NSString, useAuxiliaryFile BOOL) bool {
	ret := (C.NSString_inst_WriteToFileAtomically(o.Ptr(), path.Ptr(), (C.BOOL)(useAuxiliaryFile))) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) WriteToFileAtomicallyEncoding(path *NSString, useAuxiliaryFile BOOL, enc NSStringEncoding, error *[]*NSError) bool {

	goSlice4 := make([]unsafe.Pointer,cap(*error))
	for i := 0; i < len(*error); i++ {
		goSlice4[i] = (*error)[i].Ptr()
	}
	ret := (C.NSString_inst_WriteToFileAtomicallyEncoding(o.Ptr(), path.Ptr(), (C.BOOL)(useAuxiliaryFile), (C.NSStringEncoding)(enc), (*unsafe.Pointer)(unsafe.Pointer(&goSlice4[0])))) != 0
	(*error) = (*error)[:cap(*error)]
	for i := 0; i < len(*error); i++ {
		if goSlice4[i] == nil {
			(*error) = (*error)[:i]
			break
		}
		if (*error)[i] == nil {
			(*error)[i] = &NSError{}
			runtime.SetFinalizer((*error)[i], func(o *NSError) {
				o.Release()
			})
		}
		(*error)[i].ptr = goSlice4[i]
	}
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) ValueForUndefinedKey(key *NSString) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_ValueForUndefinedKey(o.Ptr(), key.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) IsLessThanOrEqualTo(object NSObject) bool {
	ret := (C.NSString_inst_IsLessThanOrEqualTo(o.Ptr(), object.Ptr())) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) StringByFoldingWithOptions(options NSStringCompareOptions, locale *NSLocale) *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_StringByFoldingWithOptions(o.Ptr(), (C.NSStringCompareOptions)(options), locale.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) LocalizedCapitalizedString() *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_LocalizedCapitalizedString(o.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) StringByAppendingPathExtension(str *NSString) *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_StringByAppendingPathExtension(o.Ptr(), str.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) Copy() *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_Copy(o.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) StringsByAppendingPaths(paths *NSArray) *NSArray {
	ret := &NSArray{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_StringsByAppendingPaths(o.Ptr(), paths.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSArray)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSArray) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) ClassForArchiver() Class {
	ret := (Class)(unsafe.Pointer(C.NSString_inst_ClassForArchiver(o.Ptr())))
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) URLResourceDataDidBecomeAvailable(sender *NSURL, newBytes *NSData)  {
	C.NSString_inst_URLResourceDataDidBecomeAvailable(o.Ptr(), sender.Ptr(), newBytes.Ptr())
	runtime.KeepAlive(o)
}

func (o *NSString) URLResourceDidFailLoadingWithReason(sender *NSURL, reason *NSString)  {
	C.NSString_inst_URLResourceDidFailLoadingWithReason(o.Ptr(), sender.Ptr(), reason.Ptr())
	runtime.KeepAlive(o)
}

func (o *NSString) ScriptingIsEqualTo(object NSObject) bool {
	ret := (C.NSString_inst_ScriptingIsEqualTo(o.Ptr(), object.Ptr())) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) ObjectSpecifier() *NSScriptObjectSpecifier {
	ret := &NSScriptObjectSpecifier{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_ObjectSpecifier(o.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSScriptObjectSpecifier)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSScriptObjectSpecifier) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) IsNotEqualTo(object NSObject) bool {
	ret := (C.NSString_inst_IsNotEqualTo(o.Ptr(), object.Ptr())) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) PrecomposedStringWithCompatibilityMapping() *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_PrecomposedStringWithCompatibilityMapping(o.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) StringByPaddingToLength(newLength NSUInteger, padString *NSString, padIndex NSUInteger) *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_StringByPaddingToLength(o.Ptr(), (C.NSUInteger)(newLength), padString.Ptr(), (C.NSUInteger)(padIndex)))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) SubstringFromIndex(from NSUInteger) *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_SubstringFromIndex(o.Ptr(), (C.NSUInteger)(from)))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) ScriptingProperties() *NSDictionary {
	ret := &NSDictionary{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_ScriptingProperties(o.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSDictionary)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSDictionary) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) IsGreaterThanOrEqualTo(object NSObject) bool {
	ret := (C.NSString_inst_IsGreaterThanOrEqualTo(o.Ptr(), object.Ptr())) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) InitWithBytes(bytes unsafe.Pointer, len_ NSUInteger, encoding NSStringEncoding) *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_InitWithBytes(o.Ptr(), unsafe.Pointer(bytes), (C.NSUInteger)(len_), (C.NSStringEncoding)(encoding)))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) DidChange(changeKind NSKeyValueChange, indexes *NSIndexSet, key *NSString)  {
	C.NSString_inst_DidChange(o.Ptr(), (C.NSKeyValueChange)(changeKind), indexes.Ptr(), key.Ptr())
	runtime.KeepAlive(o)
}

func (o *NSString) ReplaceValueAtIndex(index NSUInteger, key *NSString, value NSObject)  {
	C.NSString_inst_ReplaceValueAtIndex(o.Ptr(), (C.NSUInteger)(index), key.Ptr(), value.Ptr())
	runtime.KeepAlive(o)
}

func (o *NSString) TakeStoredValue(value NSObject, key *NSString)  {
	C.NSString_inst_TakeStoredValue(o.Ptr(), value.Ptr(), key.Ptr())
	runtime.KeepAlive(o)
}

func (o *NSString) SetNilValueForKey(key *NSString)  {
	C.NSString_inst_SetNilValueForKey(o.Ptr(), key.Ptr())
	runtime.KeepAlive(o)
}

func (o *NSString) Hash() NSUInteger {
	ret := (NSUInteger)(C.NSString_inst_Hash(o.Ptr()))
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) LocalizedCaseInsensitiveCompare(string *NSString) NSComparisonResult {
	ret := (NSComparisonResult)(C.NSString_inst_LocalizedCaseInsensitiveCompare(o.Ptr(), string.Ptr()))
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) MutableSetValueForKeyPath(keyPath *NSString) *NSMutableSet {
	ret := &NSMutableSet{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_MutableSetValueForKeyPath(o.Ptr(), keyPath.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSMutableSet)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSMutableSet) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) IsGreaterThan(object NSObject) bool {
	ret := (C.NSString_inst_IsGreaterThan(o.Ptr(), object.Ptr())) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) IsAbsolutePath() bool {
	ret := (C.NSString_inst_IsAbsolutePath(o.Ptr())) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) WriteToURLAtomically(url *NSURL, atomically BOOL) bool {
	ret := (C.NSString_inst_WriteToURLAtomically(o.Ptr(), url.Ptr(), (C.BOOL)(atomically))) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) WriteToURLAtomicallyEncoding(url *NSURL, useAuxiliaryFile BOOL, enc NSStringEncoding, error *[]*NSError) bool {

	goSlice4 := make([]unsafe.Pointer,cap(*error))
	for i := 0; i < len(*error); i++ {
		goSlice4[i] = (*error)[i].Ptr()
	}
	ret := (C.NSString_inst_WriteToURLAtomicallyEncoding(o.Ptr(), url.Ptr(), (C.BOOL)(useAuxiliaryFile), (C.NSStringEncoding)(enc), (*unsafe.Pointer)(unsafe.Pointer(&goSlice4[0])))) != 0
	(*error) = (*error)[:cap(*error)]
	for i := 0; i < len(*error); i++ {
		if goSlice4[i] == nil {
			(*error) = (*error)[:i]
			break
		}
		if (*error)[i] == nil {
			(*error)[i] = &NSError{}
			runtime.SetFinalizer((*error)[i], func(o *NSError) {
				o.Release()
			})
		}
		(*error)[i].ptr = goSlice4[i]
	}
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) SetObservationInfo(observationInfo unsafe.Pointer)  {
	C.NSString_inst_SetObservationInfo(o.Ptr(), unsafe.Pointer(observationInfo))
	runtime.KeepAlive(o)
}

func (o *NSString) DebugDescription() *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_DebugDescription(o.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) StoredValueForKey(key *NSString) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_StoredValueForKey(o.Ptr(), key.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) SetValueForKey(value NSObject, key *NSString)  {
	C.NSString_inst_SetValueForKey(o.Ptr(), value.Ptr(), key.Ptr())
	runtime.KeepAlive(o)
}

func (o *NSString) SetValueForKeyPath(value NSObject, keyPath *NSString)  {
	C.NSString_inst_SetValueForKeyPath(o.Ptr(), value.Ptr(), keyPath.Ptr())
	runtime.KeepAlive(o)
}

func (o *NSString) SetValueForUndefinedKey(value NSObject, key *NSString)  {
	C.NSString_inst_SetValueForUndefinedKey(o.Ptr(), value.Ptr(), key.Ptr())
	runtime.KeepAlive(o)
}

func (o *NSString) DoesContain(object NSObject) bool {
	ret := (C.NSString_inst_DoesContain(o.Ptr(), object.Ptr())) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) HasPrefix(str *NSString) bool {
	ret := (C.NSString_inst_HasPrefix(o.Ptr(), str.Ptr())) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) SubstringWithRange(range_ NSRange) *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_SubstringWithRange(o.Ptr(), (C.NSRange)(range_)))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) InitWithCStringNoCopy(bytes *Char, length NSUInteger, freeBuffer BOOL) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_InitWithCStringNoCopy(o.Ptr(), unsafe.Pointer(bytes), (C.NSUInteger)(length), (C.BOOL)(freeBuffer)))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) ValidateValueForKey(ioValue *[]*Id, inKey *NSString, outError *[]*NSError) bool {

	goSlice1 := make([]unsafe.Pointer,cap(*ioValue))
	for i := 0; i < len(*ioValue); i++ {
		goSlice1[i] = (*ioValue)[i].Ptr()
	}

	goSlice3 := make([]unsafe.Pointer,cap(*outError))
	for i := 0; i < len(*outError); i++ {
		goSlice3[i] = (*outError)[i].Ptr()
	}
	ret := (C.NSString_inst_ValidateValueForKey(o.Ptr(), (*unsafe.Pointer)(unsafe.Pointer(&goSlice1[0])), inKey.Ptr(), (*unsafe.Pointer)(unsafe.Pointer(&goSlice3[0])))) != 0
	(*ioValue) = (*ioValue)[:cap(*ioValue)]
	for i := 0; i < len(*ioValue); i++ {
		if goSlice1[i] == nil {
			(*ioValue) = (*ioValue)[:i]
			break
		}
		if (*ioValue)[i] == nil {
			(*ioValue)[i] = &Id{}
			runtime.SetFinalizer((*ioValue)[i], func(o *Id) {
				o.Release()
			})
		}
		(*ioValue)[i].ptr = goSlice1[i]
	}
	(*outError) = (*outError)[:cap(*outError)]
	for i := 0; i < len(*outError); i++ {
		if goSlice3[i] == nil {
			(*outError) = (*outError)[:i]
			break
		}
		if (*outError)[i] == nil {
			(*outError)[i] = &NSError{}
			runtime.SetFinalizer((*outError)[i], func(o *NSError) {
				o.Release()
			})
		}
		(*outError)[i].ptr = goSlice3[i]
	}
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) ValidateValueForKeyPath(ioValue *[]*Id, inKeyPath *NSString, outError *[]*NSError) bool {

	goSlice1 := make([]unsafe.Pointer,cap(*ioValue))
	for i := 0; i < len(*ioValue); i++ {
		goSlice1[i] = (*ioValue)[i].Ptr()
	}

	goSlice3 := make([]unsafe.Pointer,cap(*outError))
	for i := 0; i < len(*outError); i++ {
		goSlice3[i] = (*outError)[i].Ptr()
	}
	ret := (C.NSString_inst_ValidateValueForKeyPath(o.Ptr(), (*unsafe.Pointer)(unsafe.Pointer(&goSlice1[0])), inKeyPath.Ptr(), (*unsafe.Pointer)(unsafe.Pointer(&goSlice3[0])))) != 0
	(*ioValue) = (*ioValue)[:cap(*ioValue)]
	for i := 0; i < len(*ioValue); i++ {
		if goSlice1[i] == nil {
			(*ioValue) = (*ioValue)[:i]
			break
		}
		if (*ioValue)[i] == nil {
			(*ioValue)[i] = &Id{}
			runtime.SetFinalizer((*ioValue)[i], func(o *Id) {
				o.Release()
			})
		}
		(*ioValue)[i].ptr = goSlice1[i]
	}
	(*outError) = (*outError)[:cap(*outError)]
	for i := 0; i < len(*outError); i++ {
		if goSlice3[i] == nil {
			(*outError) = (*outError)[:i]
			break
		}
		if (*outError)[i] == nil {
			(*outError)[i] = &NSError{}
			runtime.SetFinalizer((*outError)[i], func(o *NSError) {
				o.Release()
			})
		}
		(*outError)[i].ptr = goSlice3[i]
	}
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) CString() *Char {
	ret := (*Char)(unsafe.Pointer(C.NSString_inst_CString(o.Ptr())))
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) IsLessThan(object NSObject) bool {
	ret := (C.NSString_inst_IsLessThan(o.Ptr(), object.Ptr())) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) DidChangeValueForKey(key *NSString)  {
	C.NSString_inst_DidChangeValueForKey(o.Ptr(), key.Ptr())
	runtime.KeepAlive(o)
}

func (o *NSString) DidChangeValueForKeyWithSetMutation(key *NSString, mutationKind NSKeyValueSetMutationKind, objects *NSSet)  {
	C.NSString_inst_DidChangeValueForKeyWithSetMutation(o.Ptr(), key.Ptr(), (C.NSKeyValueSetMutationKind)(mutationKind), objects.Ptr())
	runtime.KeepAlive(o)
}

func (o *NSString) CoerceValue(value NSObject, key *NSString) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_CoerceValue(o.Ptr(), value.Ptr(), key.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) StringByAbbreviatingWithTildeInPath() *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_StringByAbbreviatingWithTildeInPath(o.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) StringByStandardizingPath() *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_StringByStandardizingPath(o.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) SetScriptingProperties(scriptingProperties *NSDictionary)  {
	C.NSString_inst_SetScriptingProperties(o.Ptr(), scriptingProperties.Ptr())
	runtime.KeepAlive(o)
}

func (o *NSString) ScriptingBeginsWith(object NSObject) bool {
	ret := (C.NSString_inst_ScriptingBeginsWith(o.Ptr(), object.Ptr())) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) RangeOfCharacterFromSet(searchSet *NSCharacterSet) NSRange {
	ret := (NSRange)(C.NSString_inst_RangeOfCharacterFromSet(o.Ptr(), searchSet.Ptr()))
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) RangeOfCharacterFromSetOptions(searchSet *NSCharacterSet, mask NSStringCompareOptions) NSRange {
	ret := (NSRange)(C.NSString_inst_RangeOfCharacterFromSetOptions(o.Ptr(), searchSet.Ptr(), (C.NSStringCompareOptions)(mask)))
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) RangeOfCharacterFromSetOptionsRange(searchSet *NSCharacterSet, mask NSStringCompareOptions, rangeOfReceiverToSearch NSRange) NSRange {
	ret := (NSRange)(C.NSString_inst_RangeOfCharacterFromSetOptionsRange(o.Ptr(), searchSet.Ptr(), (C.NSStringCompareOptions)(mask), (C.NSRange)(rangeOfReceiverToSearch)))
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) CapitalizedString() *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_CapitalizedString(o.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) StringByReplacingCharactersInRange(range_ NSRange, replacement *NSString) *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_StringByReplacingCharactersInRange(o.Ptr(), (C.NSRange)(range_), replacement.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) RemoveObserverForKeyPath(observer NSObject, keyPath *NSString)  {
	C.NSString_inst_RemoveObserverForKeyPath(o.Ptr(), observer.Ptr(), keyPath.Ptr())
	runtime.KeepAlive(o)
}

func (o *NSString) RemoveObserverForKeyPathContext(observer NSObject, keyPath *NSString, context unsafe.Pointer)  {
	C.NSString_inst_RemoveObserverForKeyPathContext(o.Ptr(), observer.Ptr(), keyPath.Ptr(), unsafe.Pointer(context))
	runtime.KeepAlive(o)
}

func (o *NSString) URLResourceDidFinishLoading(sender *NSURL)  {
	C.NSString_inst_URLResourceDidFinishLoading(o.Ptr(), sender.Ptr())
	runtime.KeepAlive(o)
}

func (o *NSString) AwakeAfterUsingCoder(aDecoder *NSCoder) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_AwakeAfterUsingCoder(o.Ptr(), aDecoder.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) MutableCopyWithZone(zone *NSZone) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_MutableCopyWithZone(o.Ptr(), unsafe.Pointer(zone)))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) PathExtension() *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_PathExtension(o.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) IntValue() Int {
	ret := (Int)(C.NSString_inst_IntValue(o.Ptr()))
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) SmallestEncoding() NSStringEncoding {
	ret := (NSStringEncoding)(C.NSString_inst_SmallestEncoding(o.Ptr()))
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) ObserveValueForKeyPath(keyPath *NSString, object NSObject, change *NSDictionary, context unsafe.Pointer)  {
	C.NSString_inst_ObserveValueForKeyPath(o.Ptr(), keyPath.Ptr(), object.Ptr(), change.Ptr(), unsafe.Pointer(context))
	runtime.KeepAlive(o)
}

func (o *NSString) TakeValuesFromDictionary(properties *NSDictionary)  {
	C.NSString_inst_TakeValuesFromDictionary(o.Ptr(), properties.Ptr())
	runtime.KeepAlive(o)
}

func (o *NSString) CopyScriptingValue(value NSObject, key *NSString, properties *NSDictionary) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_CopyScriptingValue(o.Ptr(), value.Ptr(), key.Ptr(), properties.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) InitWithCharactersNoCopy(characters *Unichar, length NSUInteger, freeBuffer BOOL) *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_InitWithCharactersNoCopy(o.Ptr(), unsafe.Pointer(characters), (C.NSUInteger)(length), (C.BOOL)(freeBuffer)))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) RangeOfComposedCharacterSequenceAtIndex(index NSUInteger) NSRange {
	ret := (NSRange)(C.NSString_inst_RangeOfComposedCharacterSequenceAtIndex(o.Ptr(), (C.NSUInteger)(index)))
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) StringByRemovingPercentEncoding() *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_StringByRemovingPercentEncoding(o.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) InverseForRelationshipKey(relationshipKey *NSString) *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_InverseForRelationshipKey(o.Ptr(), relationshipKey.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) ToOneRelationshipKeys() *NSArray {
	ret := &NSArray{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_ToOneRelationshipKeys(o.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSArray)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSArray) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) IsProxy() bool {
	ret := (C.NSString_inst_IsProxy(o.Ptr())) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) CommonPrefixWithString(str *NSString, mask NSStringCompareOptions) *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_CommonPrefixWithString(o.Ptr(), str.Ptr(), (C.NSStringCompareOptions)(mask)))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) CommonPrefixWithGoString(str string, mask NSStringCompareOptions) *NSString {
	str_chr := CharWithGoString(str)
	defer str_chr.Free()
	ret := o.CommonPrefixWithString(NSStringWithUTF8String(str_chr), mask)
	return ret
}

func (o *NSString) StringByApplyingTransform(transform NSStringTransform, reverse BOOL) *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_StringByApplyingTransform(o.Ptr(), transform.Ptr(), (C.BOOL)(reverse)))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) ContainsString(str *NSString) bool {
	ret := (C.NSString_inst_ContainsString(o.Ptr(), str.Ptr())) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) LocalizedLowercaseString() *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_LocalizedLowercaseString(o.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) Init() *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_Init(o.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) MutableSetValueForKey(key *NSString) *NSMutableSet {
	ret := &NSMutableSet{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_MutableSetValueForKey(o.Ptr(), key.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSMutableSet)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSMutableSet) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) ScriptingContains(object NSObject) bool {
	ret := (C.NSString_inst_ScriptingContains(o.Ptr(), object.Ptr())) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) InitWithCharacters(characters *Unichar, length NSUInteger) *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_InitWithCharacters(o.Ptr(), unsafe.Pointer(characters), (C.NSUInteger)(length)))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) CStringUsingEncoding(encoding NSStringEncoding) *Char {
	ret := (*Char)(unsafe.Pointer(C.NSString_inst_CStringUsingEncoding(o.Ptr(), (C.NSStringEncoding)(encoding))))
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) InsertValueInPropertyWithKey(value NSObject, key *NSString)  {
	C.NSString_inst_InsertValueInPropertyWithKey(o.Ptr(), value.Ptr(), key.Ptr())
	runtime.KeepAlive(o)
}

func (o *NSString) InsertValueAtIndex(value NSObject, index NSUInteger, key *NSString)  {
	C.NSString_inst_InsertValueAtIndex(o.Ptr(), value.Ptr(), (C.NSUInteger)(index), key.Ptr())
	runtime.KeepAlive(o)
}

func (o *NSString) MutableArrayValueForKeyPath(keyPath *NSString) *NSMutableArray {
	ret := &NSMutableArray{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_MutableArrayValueForKeyPath(o.Ptr(), keyPath.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSMutableArray)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSMutableArray) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) GetClass() Class {
	ret := (Class)(unsafe.Pointer(C.NSString_inst_Class(o.Ptr())))
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) CanBeConvertedToEncoding(encoding NSStringEncoding) bool {
	ret := (C.NSString_inst_CanBeConvertedToEncoding(o.Ptr(), (C.NSStringEncoding)(encoding))) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) DecomposedStringWithCompatibilityMapping() *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_DecomposedStringWithCompatibilityMapping(o.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) Length() NSUInteger {
	ret := (NSUInteger)(C.NSString_inst_Length(o.Ptr()))
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) ReplacementObjectForKeyedArchiver(archiver *NSKeyedArchiver) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_ReplacementObjectForKeyedArchiver(o.Ptr(), archiver.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) Autorelease() *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_Autorelease(o.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) ScriptingIsLessThanOrEqualTo(object NSObject) bool {
	ret := (C.NSString_inst_ScriptingIsLessThanOrEqualTo(o.Ptr(), object.Ptr())) != 0
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) InitWithString(aString *NSString) *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_InitWithString(o.Ptr(), aString.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) InitWithGoString(aString string) *NSString {
	aString_chr := CharWithGoString(aString)
	defer aString_chr.Free()
	ret := o.InitWithString(NSStringWithUTF8String(aString_chr))
	return ret
}

func (o *NSString) LastPathComponent() *NSString {
	ret := &NSString{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_LastPathComponent(o.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*NSString)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *NSString) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func (o *NSString) ReplacementObjectForCoder(aCoder *NSCoder) *Id {
	ret := &Id{}
	ret.ptr = unsafe.Pointer(C.NSString_inst_ReplacementObjectForCoder(o.Ptr(), aCoder.Ptr()))
	if ret.ptr == nil { runtime.KeepAlive(o); return ret }
	if ret.ptr == o.ptr { runtime.KeepAlive(o); return (*Id)(unsafe.Pointer(o)) }
	runtime.SetFinalizer(ret, func(o *Id) {
		o.Release()
	})
	runtime.KeepAlive(o)
	return ret
}

func C1Alloc() *C1 {
	ret := &C1{}
	ret.ptr = unsafe.Pointer(C.c1Alloc())
	if ret.ptr == nil { return ret }
	runtime.SetFinalizer(ret,func(o *C1) {
		o.Release()
	})
	return ret
}

func (o *C1) GC() {
	if o.ptr == nil { return }
	runtime.SetFinalizer(o,func(o *C1) {
		o.Release()
	})
}

type C1Dispatch struct {
	Dealloc func(C1, C1Supermethods)
	Release func(C1, C1Supermethods)
}
var C1Lookup = map[unsafe.Pointer]C1Dispatch{}
var C1Mux sync.RWMutex

type C1Supermethods struct {
	Dealloc func()
	Release func()

}
	
func (d C1) DeallocCallback(f func(C1, C1Supermethods)) {
	C1Mux.Lock()
	dispatch := C1Lookup[d.Ptr()]
	dispatch.Dealloc = f
	C1Lookup[d.Ptr()] = dispatch
	C1Mux.Unlock()
}

func (o *C1) SuperDealloc()  {
	C.c1_super_Dealloc(o.Ptr())
	runtime.KeepAlive(o)
}

func (d C1) ReleaseCallback(f func(C1, C1Supermethods)) {
	C1Mux.Lock()
	dispatch := C1Lookup[d.Ptr()]
	dispatch.Release = f
	C1Lookup[d.Ptr()] = dispatch
	C1Mux.Unlock()
}

func (o *C1) SuperRelease()  {
	C.c1_super_Release(o.Ptr())
	runtime.KeepAlive(o)
}

func (o *C1) Dealloc()  {
	C.c1_inst_Dealloc(o.Ptr())
	runtime.KeepAlive(o)
}

func (o *C1) Release()  {
	C.c1_inst_Release(o.Ptr())
	runtime.KeepAlive(o)
}

func C2Alloc() *C2 {
	ret := &C2{}
	ret.ptr = unsafe.Pointer(C.c2Alloc())
	if ret.ptr == nil { return ret }
	runtime.SetFinalizer(ret,func(o *C2) {
		o.Release()
	})
	return ret
}

func (o *C2) GC() {
	if o.ptr == nil { return }
	runtime.SetFinalizer(o,func(o *C2) {
		o.Release()
	})
}

type C2Dispatch struct {
	MyMethod func(C2)
}
var C2Lookup = map[unsafe.Pointer]C2Dispatch{}
var C2Mux sync.RWMutex

func (d C2) MyMethodCallback(f func(C2)) {
	C2Mux.Lock()
	dispatch := C2Lookup[d.Ptr()]
	dispatch.MyMethod = f
	C2Lookup[d.Ptr()] = dispatch
	C2Mux.Unlock()
}
